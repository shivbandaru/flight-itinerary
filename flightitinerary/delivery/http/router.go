package http

import (
	"bytes"
	"encoding/json"
	"github.com/flight-itinerary/flightitinerary"
	"io"

	"github.com/labstack/echo"
)

//APIHandler struct
type APIHandler struct {
	UseCase flightitinerary.Usecase
}

//NewAPIHandler ...
func NewAPIHandler(e *echo.Echo, handler *APIHandler) {
	e.GET("/healthcheck", handler.HealthCheck)
	e.POST("/calculate", handler.CalculateItinerary)
}

// DecodeJSON ...
func DecodeJSON(data io.Reader, out interface{}) error {
	if data == nil {
		return io.EOF
	}

	decoder := json.NewDecoder(data)
	decoder.DisallowUnknownFields()
	return decoder.Decode(&out)
}

// UnmarshalJSON ...
func UnmarshalJSON(data []byte, out interface{}) error {
	return DecodeJSON(bytes.NewReader(data), out)
}
