package concurrency_patterns

import "testing"

var pipelineTests = []struct{
	Values []int
	PipelineOutput []int
}{
	{
		[]int{1,2,3,4},
		[]int{6,10, 14,18},
	},
	{
		[]int{9,11,13,15},
		[]int{38, 46, 54, 62},
	},
}

func TestPipelineGenerationFunc(t *testing.T) {

	for _, pipelineCase := range pipelineTests{
		pipelineOutput := PipelineGenerationFunc(pipelineCase.Values)
		t.Logf("%v", pipelineOutput)
		for index, output := range pipelineOutput{
			if output != pipelineCase.PipelineOutput[index] {
				t.Errorf("Expected value was : %d but got %d", pipelineCase.PipelineOutput[index], output)
			}
		}
	}

}

func TestFunctionRepaterFunc(t *testing.T) {
	FunctionRepaterFunc()

	t.Logf("Func Repeater ends")
}
