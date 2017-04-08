package interfaces

import "testing"

var minerTests = []struct {
	intArr []int
	min    int
}{
	{[]int{13, 1, 5, 15, 19, 2, 3}, 1},
	{[]int{13, 5, 1, 15, 19, 2, 0}, 0},
}

func TestMinerIntFunctionality(t *testing.T) {
	for _, minerTest := range minerTests {
		values := IntElement(minerTest.intArr)
		min := Min(values)
		if minerTest.min != min {
			t.Errorf("Expected the min element to be %v but return was %v", minerTest.min, min)
		}
	}
}

var minerFloatTests = []struct {
	intArr []float32
	min    float32
}{
	{[]float32{13.1, 1.9, 5.99, 15.9987, 19.98, 2.1, 3.09}, 1.9},
}

func TestMinerFloatFunctionality(t *testing.T) {
	for _, minerTest := range minerFloatTests {
		values := FloatElement(minerTest.intArr)
		min := Min(values)
		if minerTest.min != min {
			t.Errorf("Expected the min element to be %v but return was %v", minerTest.min, min)
		}
	}
}
