package board

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"golang.makers.tech/arrivals/flights"
)

// This creates a struct of the formatted API data from the flights package, essentially using the FlightsDataAPI struct itself
type board struct {
	flights flights.FlightDataAPI
}

// The function below is required to access data, in the format required for the functions in this package, from the main package
func NewBoard(allFlightData flights.FlightDataAPI) *board {
	return &board{
		flights: allFlightData,
	}
}

func (flightsBoard board) DisplayAPIData() table.Table {
	tbl := createTable()

	for _, flight := range flightsBoard.flights.Flights {
		addRowsToTable(flight, tbl)
	}
	return tbl
}

func (flightsBoard board) DisplayAPIDataWithFilter(airportCode string) table.Table {
	tbl := createTable()

	for _, flight := range flightsBoard.flights.Flights {
		if airportCode == flight.Origin.Code { // If the three digit airport code (typed into the terminal, eg, "go run main.go BCN" matches the origin airport code in the data the rows will be added to the table output to the terminal
			addRowsToTable(flight, tbl)
		}
	}
	return tbl
}

func createTable() table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Time", "Code", "From", "Status")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	return tbl
}

func addRowsToTable(flight flights.FlightsAPI, tbl table.Table) {
	dueTime, _ := formatTime(flight.DueTime)
	arrTime, _ := formatTime(flight.ArrTime)
	expTime, _ := formatTime(flight.ExpTime)
	flightCode := formatFlightCode(flight.Airline, flight.FlightNum)
	red := color.New(color.FgRed).SprintFunc()

	if flight.Cancelled {
		tbl.AddRow(dueTime, flightCode, flight.Origin.Code, red("Cancelled"))
	} else if checkTime(flight.ArrTime) {
		tbl.AddRow(dueTime, flightCode, flight.Origin.Code, "Landed "+arrTime)
	} else {
		tbl.AddRow(dueTime, flightCode, flight.Origin.Code, "Expected "+expTime)
	}
}

// Formats the time data from the API from a string to a time.Time formatted string
func formatTime(input string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return "", fmt.Errorf("error parsing time: %w", err)
	}
	return parsedTime.Format("15:04"), nil
}

func formatFlightCode(airline string, number string) string {
	return fmt.Sprintf("%v %v", airline, number)
}

// Checks that the time from the API data is not 'null'
func checkTime(input string) bool {
	_, err := time.Parse(time.RFC3339, input)
	return err == nil
}
