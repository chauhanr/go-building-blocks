package structures

import "testing"

var persons = []struct {
	firstName     string
	lastName      string
	expectedFName string
	expectedLName string
}{
	{"Ritesh", "Chauhan", "RITESH", "CHAUHAN"},
	{"Nitin", "Rawat", "NITIN", "RAWAT"},
}

func TestPersonToUpperFunc(t *testing.T) {

	for _, person := range persons {
		p := Person{person.firstName, person.lastName}
		p.ToUpper()
		if p.firstName != person.expectedFName || p.lastName != person.expectedLName {
			t.Errorf("Expected values (%s, %s) but received (%s, %s)", person.expectedFName, person.expectedLName, p.firstName, p.lastName)
		}
	}
}
