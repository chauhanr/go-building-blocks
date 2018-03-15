package operations

import "testing"

func TestBitwiseOperator(t *testing.T) {
	testAnd := []struct {
		A      int
		result int
	}{
		{3, 1},
		{2, 2},
		{1, 1},
		{4, 4},
		{5, 1},
		{6, 2},
		{8, 8},
		{11, 1},
		{12, 4},
		{16, 16},
		{28, 4},
		{14, 2},
	}

	/**
	  bit wise operation of a number with its negative value gives the largest power of 2 that divides the number
	  if the number is odd the value is always 1.
	*/
	for _, tc := range testAnd {
		tb := -1 * tc.A
		op := BitwiseOperator{tc.A, tb}
		if op.And() != tc.result {
			t.Errorf("bitwise operator And (%d, %d) expected : %d but got %d", tc.A, tb, tc.result, op.And())
		}
	}

	testXor := []struct {
		A      int
		B      int
		result int
	}{
		{5, 4, 1},
		{6, 4, 2},
		{6, -6, -4},
		{7, 5, 2},
		{11, 10, 1},
		{13, 12, 1},
	}

	for _, tc := range testXor {
		op := BitwiseOperator{tc.A, tc.B}
		r := op.Xor()
		if r != tc.result {
			t.Errorf("Xor operations for (%d, %d) expected %d got %d", tc.A, tc.B, tc.result, r)
		}
	}

	testNot := []struct {
		A    int
		B    int
		notA int
		notB int
	}{
		{5, 4, -6, -5},
		{6, 4, -7, -5},
	}

	for _, tc := range testNot {
		op := BitwiseOperator{tc.A, tc.B}
		na, nb := op.Not()
		if na != tc.notA && nb != tc.notB {
			t.Errorf("Not operations for (%d, %d) expected (%d, %d) got (%d, %d)", tc.A, tc.B, tc.notA, tc.notB, na, nb)
		}
	}

}
