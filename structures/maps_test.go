package structures

import (
	"fmt"
	"testing"
)

func TestMapsFunctionality(t *testing.T) {
	/**
		  map is declared as var map_name map[keytype]valuetype if we do not use new or make to instantiate a map it value is nil
	    maps. Default size after initialized map is always zero.
	*/
	var m map[string]string
	if m != nil {
		t.Errorf("The map variable must be nil when initialized.")
	}

	key := "key"
	var mapString = make(map[string]string)
	value := mapString[key]
	if value != "" {
		t.Errorf("Return type : String, The value associated with key must be nil but was %s", value)
	}
	var mapInt = make(map[string]int, 100)
	fmt.Printf("Size of the mapInt initialized as 100 is : %d \n", len(mapInt))
	valueInt := mapInt[key]
	if valueInt != 0 {
		t.Errorf("Return type : int, The value associated with key must be 0 but was %d", valueInt)
	}
}

var sampleMap = map[string]string{
	"key1": "value1",
	"key2": "value2",
	"key3": "value3",
}

// checking for the presence of a key in a map pattern.
func TestKeyPresencePattern(t *testing.T) {
	// first we log all the entires in the map The pattern here
	// shows how the range construct returns both key and values.
	for mKey, mValue := range sampleMap {
		t.Logf("[ %s : %s ]\n", mKey, mValue)
	}
	key := "ramdomkey"
	if _, isPresent := sampleMap[key]; isPresent {
		t.Errorf("Key %s should not present in the map instantiated.", key)
	}
	if _, isPresent := sampleMap["key1"]; isPresent {
		t.Logf("The value of key1 found and it has value: %s", sampleMap["key1"])
	}
}

// test case to test the deletion form the map.
func TestMapEntryDeletion(t *testing.T) {
	key := "key1"
	delete(sampleMap, key)
	if _, isPresent := sampleMap[key]; isPresent {
		t.Errorf("Key %s should have been deleted but was not.", key)
	} else {
		t.Logf("Key %s was successfully deleted from the sample map.", key)
	}
}

// first we will try to make a splice of a map pattern that can be used later when coding. s
func TestSpliceOfMap(t *testing.T) {
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	t.Logf("%v", items)
}
