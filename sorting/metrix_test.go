package sorting

import "testing"

var metriciesCases = []struct{
	metrix [][]int
	transposed [][]int
	size int
}{
	{[][]int{
		{0,1,2,3},
		{4,5,6,7},
		{8,9,10,11},
		{12,13,14,15},
	},
	    [][]int{
			{0,4,8,12},
			{1,5,9,13},
			{2,6,10,14},
			{3,7,11,15},
		},
		4,
	},

}

func TestMetrixTranspose(t *testing.T){

	for _, testCase := range metriciesCases {
		  inputMetrix := testCase.metrix
		  err := TransposeMatrix(inputMetrix)
		  if err != nil{
		  	   t.Errorf("Error in transposing the metrix %s", err.Error())
		  }else {
		  	var i, j int
			for i=0; i<	testCase.size; i++ {
				for j = 0; j < testCase.size; j++ {
					if inputMetrix[i][j] != testCase.transposed[i][j]{
						t.Errorf("The expected value at [%d, %d] was %d but was %d ", i,j,testCase.transposed[i][j], inputMetrix[i][j])
					}
				}
			}
			 t.Logf("%v", inputMetrix)
		  }
	}
}
