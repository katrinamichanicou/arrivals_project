package flights

import (
	"testing"
	"time"

	"golang.makers.tech/arrivals/test_utils"
)

func TestToStringReturnsFlightDetails(t *testing.T) {
	testFlight := Flight{Code: "BA 341", Origin: "London", DueTime: time.Date(2024, time.March, 31, 9, 30, 0, 0, time.UTC)}
	result, expected := testFlight.ToString(), "Flight BA 341 from London is expected at 09:30"
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}

func TestReadFromJSONParsesJSONFile(t *testing.T) {
	recording := test_utils.StartRecording()
	ReadFromJSON("./flightData.json")
	result := test_utils.EndRecording(recording)
	expected := ""
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}
