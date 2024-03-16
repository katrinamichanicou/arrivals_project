package flights

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type FlightDataAPI struct {
	Flights []FlightsAPI `json:"arrivals"`
}

type FlightsAPI struct {
	Origin    Origin `json:"origin"`
	Airline   string `json:"operator_iata"`
	FlightNum string `json:"flight_number"`
	Cancelled bool   `json:"cancelled"`
	DueTime   string `json:"scheduled_in"`
	ExpTime   string `json:"estimated_in"`
	ArrTime   string `json:"actual_in"`
}

type Origin struct {
	Code string `json:"code"`
}

type NullableTime struct {
	Time  *time.Time
	Valid bool
}

func (flight FlightsAPI) ParseTime(input string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return parsedTime
}

func (flight FlightsAPI) CheckTime(input string) bool {
	_, err := time.Parse(time.RFC3339, input)
	return err == nil
}

func HandleBlankTime(input string) NullableTime {
	// str := strings.Trim(string(input), `"`)
	fmt.Println(input)
	if input == "" {
		return NullableTime{Valid: false}
	}

	parsedTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return NullableTime{Valid: false}
	}

	return NullableTime{
		Time:  &parsedTime,
		Valid: true,
	}
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
	// fmt.Println("Response Body:", string(byteValue))
	json.Unmarshal(byteValue, &flightData)

	// if err := json.NewDecoder(resp.Body).Decode(&flightData); err != nil {
	// 	return FlightDataJSON{}, fmt.Errorf("error decoding JSON: %v", err)
	// }
	// fmt.Println(flightData)
	// fmt.Println("Response Status:", resp.Status)
	// fmt.Println("Response Headers:", resp.Header)

	return flightData, nil
}
