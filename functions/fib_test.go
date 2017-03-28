package functions

import "testing"

// test cases that the fib function must pass these are the first 10 value on the series.
var tests = []uint64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

func TestFibFunc(t *testing.T) {
	fn := FibFunc()
	for i, val := range tests {
		if v := fn(); val != v {
			t.Fatalf("at index %d, expected %d, got %d.", i, val, v)
		}
	}
}

// this will benchmark the operations of the function.
func BenchmarkFibFunc(b *testing.B) {
	fn := FibFunc()
	for i := 0; i < b.N; i++ {
		_ = fn()
	}
}
