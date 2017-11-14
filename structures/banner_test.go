package structures

import "testing"

func TestPrintBanner(t *testing.T) {

	arr, _ := PrintBanner("c")
	t.Logf("%v", arr)
}
