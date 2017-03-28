package functions

// FibFunc method is on which returns a fnction that returns the next element in the fibbonarchi series.
func FibFunc() func() uint64 {
	var a, b uint64 = 0, 1
	return func() uint64 {
		a, b = b, a+b
		return a
	}
}
