package flightitinerary

import (
	"context"
	"github/flight-itinerary/models"
)

type Usecase interface {
	CalculateItinerary(c context.Context, request *models.RequestJson) (*models.ResponseJson, error)
}
