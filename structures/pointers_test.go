package structures

import "testing"

var swapTests = []struct {
	a, b int64
}{
	{32, 31},
	{12, 11},
	{19, 3},
	{14, 91},
}

func TestPointers(t *testing.T) {
	for _, st := range swapTests {
		a := st.a
		b := st.b
		if SwapWithoutPointers(st.a, st.b); st.a != a && st.b != b {
			t.Errorf("Values (%d,%d) sent to swap expected a = %d, b = %d but was %d, %d respectively \n", a, b, a, b, st.a, st.b)
		}
		if SwapWithPointers(&st.a, &st.b); st.a != a && st.b != b {
			t.Errorf("Values (%d,%d) sent to swap expected a = %d, b = %d but was %d, %d respectively \n", a, b, b, a, st.a, st.b)
		}
	}
}
