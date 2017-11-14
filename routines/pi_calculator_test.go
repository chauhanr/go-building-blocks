package routines

import (
	"testing"
	"time"
)

var piResults = []struct {
	timeSeconds int
	result      float64
}{
	{1, 3.141592},
}

func TestPiCalcualtor(t *testing.T) {
	for _, piResult := range piResults {
		piValue := calculatePi(piResult.timeSeconds)
		t.Logf("Expected value of pi is %f in %d secs got %f instead", piResult.result, piResult.timeSeconds, piValue)
	}
}

func calculatePi(interval int) float64 {
	var pi float64

	pi = 0.0
	channel := make(chan float64)
	go generateSeriesSeq(channel)
	go calculateResult(channel, &pi)

	time.Sleep(1 * time.Second)

	return pi
}
