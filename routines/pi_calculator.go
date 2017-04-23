package routines

import "math"

// this series will keep generating the sequence infinitely
func generateSeriesSeq(ch chan float64) {
	for i := 0.0; ; i++ {
		ch <- i
	}
}

// this funct will change the result after each element in the series if calculated
// so the longer the routines run the more accurate the result.
func calculateResult(ch chan float64, result *float64) {
	for {
		series := <-ch
		piElem := 4 * (math.Pow(-1, series) / (2*series + 1))
		*result += piElem
	}
}
