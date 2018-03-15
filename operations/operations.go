package operations

type BitwiseOperator struct {
	A int
	B int
}

func (b *BitwiseOperator) And() int {
	return b.A & b.B
}

func (b *BitwiseOperator) Or() int {
	return b.A | b.B
}

func (b *BitwiseOperator) Xor() int {
	return b.A ^ b.B
}

func (b *BitwiseOperator) Not() (int, int) {
	return ^b.A, ^b.B
}
