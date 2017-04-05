package structures

import (
	"reflect"
	"testing"
)

func TestTagTypeStruct(t *testing.T) {
	tt := TagType{true, "Ritesh Chauhan", 1}
	for i := 0; i < 3; i++ {
		ttType := reflect.TypeOf(tt)
		nField := ttType.Field(i)
		t.Logf("Tag for field %d is %v \n", i, nField.Tag)
	}
}
