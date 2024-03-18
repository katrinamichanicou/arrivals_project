package flights

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type FlightDataAPI struct {
	Flights []FlightsAPI `json:"arrivals"`
}

type FlightsAPI struct {
	Origin    Origin `json:"origin"`
	Airline   string `json:"operator_iata"`
	FlightNum string `json:"flight_number"`
	Cancelled bool   `json:"cancelled"`
	DueTime   string `json:"scheduled_on"`
	ExpTime   string `json:"estimated_on"`
	ArrTime   string `json:"actual_on"`
}

type Origin struct {
	Code string `json:"code_iata"`
}

func ReadJSONFromOnlineAPI() (FlightDataAPI, error) {
	url := "https://aeroapi.flightaware.com/aeroapi/airports/LHR/flights/arrivals"
	password := os.Getenv("API_KEY")
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return FlightDataAPI{}, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("x-apikey", password)

	resp, err := client.Do(req)
	if err != nil {
		return FlightDataAPI{}, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	byteValue, _ := io.ReadAll(resp.Body)
	var flightData FlightDataAPI
	json.Unmarshal(byteValue, &flightData)

	return flightData, nil
}
