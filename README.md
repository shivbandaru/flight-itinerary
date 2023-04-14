# API : Determine the flight path of a person (for a given list of routes)

## API Endpoints

## GET
_____________________________________________
### /healthcheck  -- Provides the Health information of the service

## POST

------------------------------------------------------------
## /calculate  -- Calculates the origin and destination for given list of routes

### Examples
#### [["SFO", "EWR"]]   ==> ["SFO" , "EWR"]
#### [["ATL", "EWR"], ["SFO","ATL"]]   ==> ["SFO" , "EWR"]
#### [["IND", "EWR"], ["SFO","ATL"],["GSO", "IND"], ["ATL","GSO"] ]   ==> ["SFO" , "EWR"]

#### Assumption: Input is never round trip

### Sample Input Request

{
"flightRoutes":[
      {"from" : "ATL",
        "to" : "EWR"
      },
      {"from" : "SFO",
        "to" : "ATL"
      }
]
}

### Sample Output Response

{
"id": "0e1cce918a220ab29e9d7d15",
"origin": "SFO",
"destination": "EWR"
}

----------------------------------------------------------------------

## Running and Operation
% go run main.go

   ____    __

/ __/___/ /  ___

/ _// __/ _ \/ _ \

/___/\__/_//_/\___/ v4.7.0

High performance, minimalist Go web framework

https://echo.labstack.com

____________________________________O/_______

                                    

⇨ http server started on :8080

## To Do
-- Add More Unit tests (right now the code coverage is 76%)

-- Add Swagger 



