package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"golang.makers.tech/arrivals/board"
	"golang.makers.tech/arrivals/flights"
)

func printTable() {
	FlightData, err := flights.ReadJSONFromOnlineAPI()
	// This fetches the data from the API and returns the FlightDataAPI struct
	// Further details in the api.go file in the flights package

	if err != nil {
		fmt.Println("Error fetching data")
	}
	// Returns an error if the API does not return data

	flightsBoard := board.NewBoard(FlightData)
	// Creates an instance of the flight board using the FlightData pulled from the API
	if len(os.Args) < 2 {
		fmt.Println("      LONDON HEATHROW ARRIVALS")
		flightsBoard.DisplayAPIData().Print()
		// Displays the flights per the DisplayAPIData() function in the board package
	} else {
		filter := os.Args[1]
		fmt.Printf(" LONDON HEATHROW ARRIVALS FROM %v \n", filter)
		flightsBoard.DisplayAPIDataWithFilter(filter).Print()
		// If an origin airport code is specified when running main.go in the terminal
		// the table is filtered for flights originating from that airport
	}
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	// cmd.Stdout = os.Stdout ensure that the output from the 'cmd' is in the same terminal where the program is running
	cmd.Run()
}

func main() {
	clearTerminal()
	printTable()
	ticker := time.NewTicker(1 * time.Minute)
	// time.NewTicker creates a new goroutine of a ticker for the duration specified in the parentheses (every minute)
	defer ticker.Stop()
	// This stops the new ticker when the main function exits

	for range ticker.C {
		// ticker.C is a channel that is used to  receive "tick" events from the time.Ticker
		// When a new time.Ticker is created (using time.NewTicker(duration)) each time the ticker's internal timer fires
		// it sends a current time value to the 'C' channel
		clearTerminal()
		go printTable()
	}

}
