package http

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github/flight-itinerary/flightitinerary/delivery/itinerary"
	"github/flight-itinerary/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateItinerary(t *testing.T) {

	useCase := itinerary.NewItineraryUsecase()
	apiHandler := &APIHandler{UseCase: useCase}
	e := echo.New()

	t.Run("should return 200 status ok", func(t *testing.T) {
		input := models.RequestJson{FlightRoutes: []models.Route{{From: "ATL", To: "EWR"}, {From: "SFO", To: "ATL"}}}

		payload, _ := json.Marshal(input)
		req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		apiHandler.CalculateItinerary(c)

		assert.Equal(t, http.StatusOK, rec.Code)

	})

	t.Run("should return 500 status ", func(t *testing.T) {
		input := models.RequestJson{FlightRoutes: []models.Route{{From: "ATL", To: "EWR"}, {From: "SFO", To: "ATL"}, {From: "EWR", To: "SFO"}}}

		payload, _ := json.Marshal(input)
		req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		apiHandler.CalculateItinerary(c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

	})
}
