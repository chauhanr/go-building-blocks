package functions

import (
	"testing"
)

var min_max_tests = []struct {
	numbers  []int64
	min, max int64
}{
	{[]int64{1, 39, 19, 15, 0, 133, 41, -1, -55}, -55, 133},
	{[]int64{1, 39, 19, 15, 0, 249, 133, 41, -1, -59, -199}, -199, 249},
}

func TestMinMaxFunc(t *testing.T) {

	for _, test := range min_max_tests {
		min, max := MinMax(test.numbers)

		if min != test.min || max != test.max {
			t.Errorf("Min and Max must be %d, %d but were %d, %d", test.min, test.max, min, max)
		}
	}

}
