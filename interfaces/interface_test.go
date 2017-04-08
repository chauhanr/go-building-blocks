package interfaces

import (
	"reflect"
	"testing"
)

var interfaces = []struct {
	s         Shaper
	area      float32
	shapeType string
}{
	{&Circle{7.0}, 153.938049, "*interfaces.Circle"},
	{&Square{4.0}, 16.0, "*interfaces.Square"},
}

func TestShaperInterfaceTypes(t *testing.T) {
	for _, shape := range interfaces {
		shaper := shape.s
		if shaper.Area() != shape.area {
			t.Errorf("Expected Area was : %f but calculated value was : %f ", shape.area, shaper.Area())
		}
		kind := reflect.TypeOf(shaper)
		// also check the type returned by the interface using reflect package
		//t.Logf("shape kind is : %s", kind.String())
		if kind.String() != shape.shapeType {
			t.Errorf("The expected value of shaper was %s but we got %s", shape.shapeType, kind.String())
		}
	}
}

var classifierTests = []struct {
	arguments []interface{}
	expected  []string
}{
	{[]interface{}{13, -14.4, "Belgrade", complex(1, 3), nil}, []string{"int", "float", "string", "undefined", "nil"}},
	{[]interface{}{false, true, "Myth", new(Square), "mint"}, []string{"bool", "bool", "string", "Square", "string"}},
}

func TestClassiferFunc(t *testing.T) {
	for _, testCase := range classifierTests {
		results := Classifier(testCase.arguments...)
		//log.Printf("Length of input args : %d \n", len(testCase.arguments))
		//log.Printf("Length of results : %d \n", len(results))
		args := testCase.expected
		if len(results) != len(args) {
			t.Errorf("The input arguments are not equal to the output results length\n")
		} else {
			for i := 0; i < len(args); i++ {
				if args[i] != results[i] {
					t.Errorf("Expected type %s but found type %s \n", args[i], results[i])
				}
			}
		}
	}
}

func TestEmptyInterface(t *testing.T) {
	var spl specialString = "Ritesh"
	results := Classifier(spl)

	if len(results) != 1 {
		t.Errorf("The input arguments are not equal to the output results length\n")
	} else {
		for i := 0; i < len(results); i++ {
			if "string" != results[i] {
				t.Errorf("Expected type %s but found type %s \n", "string", results[i])
			}
		}
	}

}
