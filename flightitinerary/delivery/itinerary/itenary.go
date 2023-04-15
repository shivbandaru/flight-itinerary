package itinerary

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/flight-itinerary/flightitinerary"
	"github.com/flight-itinerary/logger"
	"github.com/flight-itinerary/models"
)

//ItineraryUsecase...
type ItineraryUsecase struct {
}

func NewItineraryUsecase() flightitinerary.Usecase {
	return &ItineraryUsecase{}
}

//CalculateItinerary calculates the origin and destination for a given list of routes
func (s *ItineraryUsecase) CalculateItinerary(c context.Context, request *models.RequestJson) (*models.ResponseJson, error) {

	requestId := c.Value(models.RequestID).(string)
	switch len(request.FlightRoutes) {
	case 0:
		return &models.ResponseJson{Id: requestId}, nil
	case 1:
		return &models.ResponseJson{Id: requestId, Origin: request.FlightRoutes[0].From, Destination: request.FlightRoutes[0].To}, nil
	default:
		routeMap := make(map[string]string)
		routeReverseMap := make(map[string]string)
		var origin, destination string
		originCount := 0
		destCount := 0

		//Creating a hashmap with from as key and  To as value
		for _, route := range request.FlightRoutes {
			routeMap[route.From] = route.To
			routeReverseMap[route.To] = route.From
		}

		// below code will get the origin airport code ,
		//logic: Source airport code cannot be in destination ariport code
		for from := range routeMap {
			if _, ok := routeReverseMap[from]; !ok {
				origin = from
				originCount++
			}
		}

		//traverse the map from start to end to get the destination
		for _, to := range routeMap {
			if _, ok := routeMap[to]; !ok {
				destination = to
				destCount++
			}
		}

		//There could be a chance if there is a missing route , then we cannot give incorrect origin and destination
		if origin == "" || destination == "" || originCount > 1 || destCount > 1 {
			logger.Log.ErrorC(c, "Failed to find origin and destination for given list of routes")
			return nil, fmt.Errorf("unable to find the origin and destination,this could be a roundtrip or missing route")
		}

		response := models.ResponseJson{Id: requestId, Origin: origin, Destination: destination}
		logger.Log.InfoC(c, "Response is ", spew.Sprintf("%+v", response))

		return &response, nil
	}

}
