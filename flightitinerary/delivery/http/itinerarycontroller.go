package http

import (
	"context"
	"github/flight-itinerary/logger"
	"github/flight-itinerary/models"
	"strconv"

	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
)

// CalculateItinerary method
// @Summary  Calculate the origin and destination for a given one-way trip
// @Produce json
// @Tags Flight-Itinerary APIs
// @Param data body models.RequestJson true "body"
// @Success 200 {object} models.ResponseJson
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /calculate [post]
func (handler *APIHandler) CalculateItinerary(c echo.Context) error {

	request := models.RequestJson{}

	if err := c.Bind(&request); err != nil {
		//if err := DecodeJSON(c.Request().Body, &request); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseError{Code: strconv.Itoa(http.StatusBadRequest), Message: err.Error()})
	}

	ctx := context.WithValue(c.Request().Context(), models.RequestID, c.Response().Header().Get("RequestID"))

	logger.Log.InfoC(ctx, "Incoming request is ", spew.Sprintf("%+v", request))

	response, err := handler.UseCase.CalculateItinerary(ctx, &request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.ResponseError{Code: strconv.Itoa(http.StatusInternalServerError), Message: err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}

// HealthCheck endpoint
// @Summary Get HealthCheck
// @Description Get HealthCheck
// @Tags HealthCheck APIs
// @Produce json
// @Success 200
// @Router /healthcheck [get]
func (handler *APIHandler) HealthCheck(c echo.Context) error {
	ctx := context.WithValue(c.Request().Context(), models.RequestID, c.Response().Header().Get("RequestID"))
	logger.Log.InfoC(ctx, "Healthy")
	return c.JSON(http.StatusOK, "OK")
}
