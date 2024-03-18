package board

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"golang.makers.tech/arrivals/flights"
)

type board struct {
	flights flights.FlightDataAPI
}

// This function is for the flight data from hardcoded JSON file
// func (flightsBoard board) Display() table.Table {
// 	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
// 	columnFmt := color.New(color.FgYellow).SprintfFunc()
// 	red := color.New(color.FgRed).SprintFunc()

// 	tbl := table.New("Time", "Code", "From", "Status")
// 	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

// 	for _, flight := range flightsBoard.flights.Flights {
// 		if flight.Status.Cancelled {
// 			tbl.AddRow(flight.DueTime.Format15_04(), flight.Code, flights.GetAirportCode(flightsBoard.flights.Airports, flight.Origin), red("Cancelled"))
// 		} else if flight.Status.ArrTime.Format15_04() != "00:00" {
// 			tbl.AddRow(flight.DueTime.Format15_04(), flight.Code, flights.GetAirportCode(flightsBoard.flights.Airports, flight.Origin), "Landed "+flight.Status.ArrTime.Format15_04())
// 		} else {
// 			tbl.AddRow(flight.DueTime.Format15_04(), flight.Code, flights.GetAirportCode(flightsBoard.flights.Airports, flight.Origin), "Expected "+flight.Status.ExpTime.Format15_04())
// 		}
// 	}

// 	return tbl
// }

// This function is for the flight data from the hardcoded JSON file
// func (flightsBoard board) DisplayWithFilter(airportCode string) table.Table {
// 	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
// 	columnFmt := color.New(color.FgYellow).SprintfFunc()
// 	red := color.New(color.FgRed).SprintFunc()

// 	tbl := table.New("Time", "Code", "From", "Status")
// 	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

// 	for _, flight := range flightsBoard.flights.Flights {
// 		if flightsBoard.flights.Airports[airportCode] == flight.Origin {
// 			if flight.Status.Cancelled {
// 				tbl.AddRow(flight.DueTime.Format15_04(), flight.Code, flights.GetAirportCode(flightsBoard.flights.Airports, flight.Origin), red("Cancelled"))
// 			} else if flight.Status.ArrTime.Format15_04() != "00:00" {
// 				tbl.AddRow(flight.DueTime.Format15_04(), flight.Code, flights.GetAirportCode(flightsBoard.flights.Airports, flight.Origin), "Landed "+flight.Status.ArrTime.Format15_04())
// 			} else {
// 				tbl.AddRow(flight.DueTime.Format15_04(), flight.Code, flights.GetAirportCode(flightsBoard.flights.Airports, flight.Origin), "Expected "+flight.Status.ExpTime.Format15_04())
// 			}
// 		}
// 	}

// 	return tbl
// }

func NewBoard(allFlightData flights.FlightDataAPI) *board {
	return &board{
		flights: allFlightData,
	}
}

func (flightsBoard board) DisplayAPIData() table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	red := color.New(color.FgRed).SprintFunc()

	tbl := table.New("Time", "Code", "From", "Status")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, flight := range flightsBoard.flights.Flights {
		if flight.Cancelled {
			tbl.AddRow(formatTime(flight.DueTime), formatFlightCode(flight.Airline, flight.FlightNum), flight.Origin.Code, red("Cancelled"))
		} else if flight.CheckTime(flight.ArrTime) {
			tbl.AddRow(formatTime(flight.DueTime), formatFlightCode(flight.Airline, flight.FlightNum), flight.Origin.Code, "Landed "+formatTime(flight.ArrTime))
		} else {
			tbl.AddRow(formatTime(flight.DueTime), formatFlightCode(flight.Airline, flight.FlightNum), flight.Origin.Code, "Expected "+formatTime(flight.ExpTime))
		}
	}
	return tbl
}

func (flightsBoard board) DisplayAPIDataWithFilter(airportCode string) table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	red := color.New(color.FgRed).SprintFunc()

	tbl := table.New("Time", "Code", "From", "Status")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, flight := range flightsBoard.flights.Flights {
		if airportCode == flight.Origin.Code {
			if flight.Cancelled {
				tbl.AddRow(formatTime(flight.DueTime), formatFlightCode(flight.Airline, flight.FlightNum), flight.Origin.Code, red("Cancelled"))
			} else if flight.CheckTime(flight.ArrTime) {
				tbl.AddRow(formatTime(flight.DueTime), formatFlightCode(flight.Airline, flight.FlightNum), flight.Origin.Code, "Landed "+formatTime(flight.ArrTime))
			} else {
				tbl.AddRow(formatTime(flight.DueTime), formatFlightCode(flight.Airline, flight.FlightNum), flight.Origin.Code, "Expected "+formatTime(flight.ExpTime))
			}
		}
	}
	return tbl
}

func formatTime(input string) string {
	parsedTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		fmt.Errorf("error: %w", err)
	}
	return parsedTime.Format("15:04")
}

func formatFlightCode(airline string, number string) string {
	return fmt.Sprintf("%v %v", airline, number)
}
