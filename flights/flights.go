package flights

import (
	"fmt"
	"time"
)

type Flight struct {
	Code    string
	Origin  string
	DueTime time.Time
}

func HelloWorld() {
	fmt.Println("Hello from the flights package!")
}

func (flightInfo Flight) ToString() string {
	return fmt.Sprintf("Flight %s from %s is expected at %s", flightInfo.Code, flightInfo.Origin, flightInfo.DueTime.Format("15:04"))
}
