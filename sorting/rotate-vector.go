package sorting

import "errors"

/**
	Programming pearls column 2 | problem 3
	This method will rotate a vector on a pivot it used reversing technique for each subsection followed by a total reverse
	1. reverse vector between 0 and pivot-1
 	2. reverse vector between pivot and end-1
	3. reverse the entire pivot.
*/

func RotateVector(vector []string, rotatePivot int) ([]string, error, int){
	traversal := 1;
	if len(vector) < rotatePivot {
		return nil, errors.New("Cannot have a rotate pivot larger than length of vector"), traversal
	}
	if rotatePivot == 0 {
		return nil, errors.New("Cannot have a rotate pivot as zero"), traversal
	}
	end := len(vector)-1
	// step 1: reverse from start to the pivot
	traversal += reverse(vector,0, rotatePivot-1)
	// step 2: reverse from pivot to end
	traversal +=reverse(vector, rotatePivot, end)
	// step 3: reverse the entire string
	traversal +=reverse(vector, 0, end)

	return vector, nil, traversal
}
// simple function that reverses the section of the array
func reverse(vector []string, start int, end int) int {
	traversal := 0
	for end > start{
		temp := vector[start]
		vector[start] = vector[end]
		vector[end] = temp
		end--
		start++
		traversal++
	}
	return traversal
}
