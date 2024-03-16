package main

// import (
// 	"fmt"
// 	"testing"
// 	"time"

// 	"golang.makers.tech/arrivals/flights"
// 	"golang.makers.tech/arrivals/test_utils"
// )

// func TestToStringReturnsMultiLine(t *testing.T) {
// 	recording := test_utils.StartRecording()
// 	testFlight1 := flights.Flight{Code: "BA 321", Origin: "New York", DueTime: time.Date(2024, time.March, 31, 9, 30, 0, 0, time.UTC)}
// 	testFlight2 := flights.Flight{Code: "BA 625", Origin: "Barcelona", DueTime: time.Date(2024, time.March, 31, 18, 55, 0, 0, time.UTC)}
// 	testFlight3 := flights.Flight{Code: "BA 828", Origin: "Athens", DueTime: time.Date(2024, time.March, 31, 22, 10, 0, 0, time.UTC)}

// 	testBoard := Board{Flights: []flights.Flight{testFlight1, testFlight2, testFlight3}}
// 	fmt.Println(testBoard.Display())
// 	result := test_utils.EndRecording(recording)
// 	expected := `Flight Arrivals
// 	09:30 New York BA 321
// 	18:55 Barcelona BA 625
// 	22:10 Athens BA 828`

// 	if result != expected {
// 		t.Errorf("result is %v but %v was expected", result, expected)
// 	}
// }
