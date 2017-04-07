package interfaces

import (
	"log"
	"testing"
)

var ints = []struct {
	intArr    []int
	sortedArr []int
}{
	{[]int{18, 9, 12, 1, 3, 19, 21}, []int{1, 3, 9, 12, 18, 19, 21}},
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
}

func TestSortForIntArray(t *testing.T) {
	for _, arr := range ints {
		input := arr.intArr
		expected := arr.sortedArr
		raw := IntArray(input)
		Sort(raw)
		log.Printf("%v", raw)
		for i := 0; i < len(raw); i++ {
			if raw[i] != expected[i] {
				t.Errorf("Expected element %d at index %d but found %d ", expected[i], i+1, raw[i])
			}
		}
	}
}

var strs = []struct {
	strArr       []string
	sortedstrArr []string
}{
	{[]string{"a", "z", "m", "n", "y", "x", "p"}, []string{"a", "m", "n", "p", "x", "y", "z"}},
	{[]string{"ritesh", "rajinder", "satish", "rajesh", "vivek", "ankur", "shorav"}, []string{"ankur", "rajesh", "rajinder", "ritesh", "satish", "shorav", "vivek"}},
}

func TestSortForStringArray(t *testing.T) {
	for _, arr := range strs {
		input := arr.strArr
		expected := arr.sortedstrArr
		raw := StringArray(input)
		if IsSorted(raw) == true {
			t.Errorf("The %v array is not sorted but IsSorted Function says it is \n", raw)
		}

		Sort(raw)
		log.Printf("%v", raw)
		for i := 0; i < len(raw); i++ {
			if raw[i] != expected[i] {
				t.Errorf("Expected element %s at index %d but found %s ", expected[i], i+1, raw[i])
			}
		}
		if IsSorted(raw) == false {
			t.Errorf("The %v array must be sorted at this stage but is not sorted \n", raw)
		}
	}
}
