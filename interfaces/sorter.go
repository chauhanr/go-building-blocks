package interfaces

// Sorter used to sort various arrays that may come in.
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Sort function is the algorithm that will sort the array in questiom
func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

// IsSorted method checks if array is already sorted.
func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

// IntArray gives us a list that we need to use when sorting integers
type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int) {
	tmp := p[i]
	p[i] = p[j]
	p[j] = tmp
}

// StringArray is teh type to handle string array interface
type StringArray []string

func (p StringArray) Len() int {
	return len(p)
}

func (p StringArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringArray) Swap(i, j int) {
	tmp := p[i]
	p[i] = p[j]
	p[j] = tmp
}
