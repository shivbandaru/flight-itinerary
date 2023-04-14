package models

type RequestJson struct {
	FlightRoutes []Route `json:"flightRoutes,required"`
}

type Route struct {
	From string `json:"from,required"`
	To   string `json:"to,required"`
}

type ResponseJson struct {
	Id          string `json:"id"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

// ResponseError model
type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
