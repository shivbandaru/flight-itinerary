package main

import (
	"context"
	"fmt"
	fi "github.com/flight-itinerary/flightitinerary/delivery/http"
	"github.com/flight-itinerary/flightitinerary/delivery/itinerary"
	"github.com/flight-itinerary/middleware"

	"github.com/labstack/echo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Flight-Itinerary Service
// @version 1.0.0
// @contact.name   Siva Bandaru
// @contact.email  bandaru.kumar.s@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8086
// @BasePath /
// @query.collection.format multi

func main() {

	//Create a signal channel to capture interrupt and terminate signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	//Crete a new echo
	e := echo.New()
	midl := middleware.InitMiddleware()
	e.Use(midl.RequestID)

	useCase := itinerary.NewItineraryUsecase()
	handler := &fi.APIHandler{UseCase: useCase}

	fi.NewAPIHandler(e, handler)

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// =========================================================================
	// Start Swagger handler

	fmt.Println("Server started on :8080")

	//wait for interrupt or termination signal
	<-signalChan

	//Create a deadline for shutting down the server gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Shutdown the server gracefully
	if err := e.Shutdown(ctx); err != nil {
		fmt.Println("Failed to gracefully shutdown the server:", err)
		os.Exit(1)
	}

	fmt.Println("server gracefully shutdown")

}
