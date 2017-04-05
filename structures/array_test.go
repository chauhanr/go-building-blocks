package structures

import (
	"sort"
	"testing"
)

var arr = []string{"new delhi", "canberra", "washingtion", "manila"}
var sortedStr = "canberra,manila,new delhi,washingtion,"

func TestArraySort(t *testing.T) {
	unsorted := ""
	for _, val := range arr {
		unsorted += val + ","
	}
	t.Logf("%s", unsorted)

	sort.Strings(arr)
	sorted := ""
	for _, val := range arr {
		sorted += val + ","
	}
	t.Logf("%s", sorted)

	if sortedStr != sorted {
		t.Errorf("Array not sorted properly should be : \n %s\n but is : %s ", sortedStr, sorted)
	}
}
