package functions

import "testing"

var mulTests = []struct {
	a, b     int64
	expected int64
}{
	{1, 1, 1},
	{2, 2, 4},
	{3, 3, 9},
	{5, 6, 30},
	{1, 3, 3},
}

// this test will test the Mul function that is simple as it takes 2 integers as inputs and then returns another integer.
func TestMul(t *testing.T) {
	for _, mt := range mulTests {
		if v := Mul(mt.a, mt.b); v != mt.expected {
			t.Errorf("Mul (%d, %d) returned %d, expected %d ", mt.a, mt.b, v, mt.expected)
		}
	}
}
