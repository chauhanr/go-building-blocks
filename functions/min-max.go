package functions

import (
	"math"
)

// function will accept an array of integers and find the min and max elements
func MinMax(numbers []int64) (min int64, max int64) {
	max = math.MinInt64
	min = math.MaxInt64
	//log.SetFlags(log.Llongfile)
	//var where = log.Print

	for _, num := range numbers {
		//where("(Min, Max) is \n", min, max)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return min, max
}
