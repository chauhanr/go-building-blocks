package structures

import "strings"

// Person struct with last and first name as the attributes.
type Person struct {
	firstName string
	lastName  string
}

// ToUpper method will convert the first and last names to upper values.
func (person *Person) ToUpper() {
	person.firstName = strings.ToUpper(person.firstName)
	person.lastName = strings.ToUpper(person.lastName)
}
