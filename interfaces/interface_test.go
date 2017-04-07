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
