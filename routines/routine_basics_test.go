package routines

import (
	"testing"
	"time"
)

type CityStruct struct {
	cities []string
}

var cityResult = []string{"New Delhi", "Colombo", "Washington", "Timphu", "Seoul"}

func TestRunGoRoutines(t *testing.T) {
	ch := make(chan string)
	cities := CityStruct{}

	go sendData(ch, cityResult)
	go getData(ch, &cities)
	time.Sleep(1e9)

	for i, city := range cities.cities {
		if city != cityResult[i] {
			t.Logf("Expected %s and received %s ", cityResult[i], city)
		}
	}
}
