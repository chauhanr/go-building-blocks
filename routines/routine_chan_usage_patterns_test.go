package routines

import (
	"testing"
	"time"
)

type CityStruct struct {
	cities []string
}

var cityResult = []string{"New Delhi", "Colombo", "Washington", "Timphu", "Seoul"}

/**
The test shows how the unbuffered channel must be used when two go routine need to communicate
1. the sender must be finite and must run in its routine. if it is not run in its own routine it will block the channel
and will cause the go runtime to throw a panic.
2. similarly the receiver function must also run in its own routine otherwise it causes the main test routine to block
and the last line is not able to run causing a deadlock too.

the final step asks the main routine to sleep allows both the routines to perform their operations otherwise we will
not get all the results.
*/
func TestUnbufferedChannelPattern(t *testing.T) {
	ch := make(chan string)
	cities := CityStruct{}

	go sendData(ch, cityResult)
	go getData(ch, &cities)
	time.Sleep(1e9)

	if cities.cities == nil {
		t.Errorf("The routines did not get the appropriate time to run.")
	}

	for i, city := range cities.cities {
		if city != cityResult[i] {
			t.Errorf("Expected %s and received %s \n", cityResult[i], city)
		} else {
			t.Logf("Capital : %s\n", city)
		}
	}
}
