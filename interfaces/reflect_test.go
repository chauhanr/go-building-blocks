package interfaces

import (
	"go-building-blocks/structures"
	"reflect"
	"testing"
)

var reflectTests = []struct {
	data      interface{}
	dataType  string
	dataValue string
	dataKind  string
}{
	{3.4, "float64", "<float64 Value>", "float64"},
}

func TestReflectPackageFuncs(t *testing.T) {

	for _, refTest := range reflectTests {
		f := refTest.data
		ty := reflect.TypeOf(f)
		if ty.String() != refTest.dataType {
			t.Errorf("Expected data type was %s but got %s ", refTest.dataType, ty.String())
		}
		v := reflect.ValueOf(f)
		if v.String() != refTest.dataValue {
			t.Errorf("Expected data value was %s but got %s ", refTest.dataValue, v.String())
		}
		if v.Kind().String() != refTest.dataKind {
			t.Errorf("Expected data kind was %s but got %s ", refTest.dataKind, v.Kind().String())
		}
		t.Logf("Value of the item is : %v\n", v.Interface())
		t.Logf("Type of the item is : %T\n", v.Interface())
	}
}

func TestChangeValueWithReflect(t *testing.T) {
	f := 3.4
	v := reflect.ValueOf(f)
	if !v.CanSet() {
		t.Logf("As value is not passed by reference we cannot set it.\n")
	} else {
		t.Errorf("Error the element should not be settable as it is passed by value")
	}

	v = reflect.ValueOf(&f)
	if !v.CanSet() {
		t.Logf("The value is not settable because the Elem() is not used.")
	}

	v = v.Elem()
	if v.CanSet() {
		t.Logf("The value is settable because Elem() has been called. ")
	}

	newf := 3.145
	v.SetFloat(newf)
	if v.Interface() != newf {
		t.Errorf("The new float value should have been %v but was %v", newf, v)
	}
}

var refStructTests = []struct {
	person          structures.Person
	interfaceType   string
	interfaceKind   string
	fieldCollection []string
}{
	{structures.Person{"Ritesh", "Chauhan", "8800439536", 37},
		"structures.Person", "struct",
		[]string{"FirstName", "LastName", "Age", "PhoneNumber"}},
}

func TestPersonStruct(t *testing.T) {
	for _, personTest := range refStructTests {
		var p interface{} = personTest.person
		refType := reflect.TypeOf(p).String()
		t.Logf("The RefType of the struct interface is : %s", refType)
		if refType != personTest.interfaceType {
			t.Errorf("Expected reflection type was %s but found %s ", refType, personTest.interfaceType)
		}
		value := reflect.ValueOf(p)
		refKind := value.Kind().String()
		t.Logf("The RefKind of the struct interface is : %s", refKind)
		if refKind != personTest.interfaceKind {
			t.Errorf("Expected reflection kind was %s but found %s ", refKind, personTest.interfaceKind)
		}
		for i := 0; i < value.NumField(); i++ {
			t.Logf("Field at %d : %v", i, value.Field(i))
		}
		t.Logf("Method count of %s is %d ", refType, value.NumMethod())
		for j := 0; j < value.NumMethod(); j++ {
			t.Logf("Method at %d : %v", j, value.Method(j))
		}
	}
}
