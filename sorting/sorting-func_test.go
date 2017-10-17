package sorting

import (
	"testing"
)

var rotateVectorTest = []struct{
	vector []string
	pivot int
	rotatedVector []string
}{
	{   []string{"a", "b", "c","d", "e", "f", "g", "h"},
		3,
		[]string{"d", "e", "f", "g", "h","a", "b", "c"},
	},
	{   []string{"a", "b", "c","d", "e", "f", "g", "h"},
		5,
		[]string{"f", "g", "h","a", "b", "c","d", "e"},
	},
	{   []string{"a", "b", "c","d", "e", "f", "g", "h"},
		7,
		[]string{"h", "a", "b", "c","d", "e", "f", "g"},
	},
	{   []string{"a", "b", "c","d", "e", "f", "g", "h"},
		1,
		[]string{ "b", "c","d", "e", "f", "g", "h", "a"},
	},
}


func TestRotateVector(t *testing.T) {

	for _, testCase := range rotateVectorTest{
		//t.Logf("Testing for vector %v",testCase.vector)
		rotatedVector, err, iteration:= RotateVector(testCase.vector, testCase.pivot)
		if err != nil{
			t.Errorf("vector should have been %v, but an error occured %s", testCase.rotatedVector, err.Error())
		}else{
			for index, actual := range rotatedVector {
				expected := testCase.rotatedVector[index]
				if actual != expected{
					t.Errorf("Vector should have value %s at %d but found %s", expected, index,actual)
				}else{
					// do nothing
				}
			}
			t.Logf("Testing for vector for pivot %d result is : %v",testCase.pivot,rotatedVector)
		}
		t.Logf("Total number of iteration is %d", iteration)
	}

}