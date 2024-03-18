package main

import (
	"fmt"
	"os"

	"golang.makers.tech/arrivals/board"
	"golang.makers.tech/arrivals/flights"
)

func main() {
	if len(os.Args) < 2 {
		// FlightData, err := flights.ReadFromJSON("flightData.json")
		// FlightData, err := flights.ReadJSONFromAPI()
		FlightData, err := flights.ReadJSONFromOnlineAPI()
		if err != nil {
			fmt.Println("JSON file not found")
		}
		flightsBoard := board.NewBoard(FlightData)

		// flightsBoard.Display().Print()
		flightsBoard.DisplayAPIData().Print()
	} else {
		filter := os.Args[1]
		// FlightData, err := flights.ReadFromJSON("flightData.json")
		// FlightData, err := flights.ReadJSONFromAPI()
		FlightData, err := flights.ReadJSONFromOnlineAPI()
		if err != nil {
			fmt.Println("JSON file not found")
		}
		flightsBoard := board.NewBoard(FlightData)

		// flightsBoard.DisplayWithFilter(filter).Print()
		flightsBoard.DisplayAPIDataWithFilter(filter).Print()
	}
}
