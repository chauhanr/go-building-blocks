package structures

// SwapWithoutPointers will be a poor example of swap because it will not swap the values.
func SwapWithoutPointers(a int64, b int64) {
	var temp int64
	temp = a
	a = b
	b = temp
}

// SwapWithPointers will be the correc implementation of the swap method where we swap the values.
func SwapWithPointers(a *int64, b *int64) {
	var temp *int64
	temp = a
	a = b
	b = temp
}
