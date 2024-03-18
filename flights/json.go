// No longer required : Superseded by api.go

package flights

import (
	"encoding/json"
	//"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type FlightDataJSON struct {
	Flights  []FlightJSON      `json:"flights"`
	Airports map[string]string `json:"airports"`
}

type FlightJSON struct {
	Origin  string     `json:"from"`
	Code    string     `json:"code"`
	DueTime myTime     `json:"scheduled_arrival"`
	Status  StatusJSON `json:"status"`
}

type StatusJSON struct {
	ArrTime   myTime `json:"arrived"`
	Cancelled bool   `json:"cancelled"`
	ExpTime   myTime `json:"expected_at"`
}

type myTime time.Time

func (t *myTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)
	if str == "" {
		*t = myTime(time.Time{})
		return nil
	}
	pt, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}
	*t = myTime(pt)
	return nil
}

func (t myTime) Format15_04() string {
	tt := time.Time(t)
	return tt.Format("15:04")
}

func ReadJSONFromAPI() (FlightDataJSON, error) {
	url := "https://go-flights-api.onrender.com/flights?airport=LHR"
	password := os.Getenv("API_PASSWORD")
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return FlightDataJSON{}, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("x-makers-password", password)

	resp, err := client.Do(req)
	if err != nil {
		return FlightDataJSON{}, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	byteValue, _ := io.ReadAll(resp.Body)
	var flightData FlightDataJSON
	fmt.Println("Response Body:", string(byteValue))
	json.Unmarshal(byteValue, &flightData)

	// if err := json.NewDecoder(resp.Body).Decode(&flightData); err != nil {
	// 	return FlightDataJSON{}, fmt.Errorf("error decoding JSON: %v", err)
	// }
	fmt.Println(flightData)
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)

	return flightData, nil
}

func GetAirportCode(airports map[string]string, origin string) string {
	for k, v := range airports {
		if v == origin {
			return k
		}
	}
	return ""
}

// ** Function to read hardcoded JSON file - not linked to an API
func ReadFromJSON(filepath string) (FlightDataJSON, error) {
	jsonFile, err := os.Open(filepath)

	if err != nil {
		return FlightDataJSON{}, fmt.Errorf("error: %v", err)
	}
	defer jsonFile.Close()

	var flightData FlightDataJSON
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &flightData)

	return flightData, nil
}
