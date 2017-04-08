package structures

import "strings"

// Person struct with last and first name as the attributes.
type Person struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
}

// ToUpper method will convert the first and last names to upper values.
func (person *Person) ToUpper() {
	person.FirstName = strings.ToUpper(person.FirstName)
	person.LastName = strings.ToUpper(person.LastName)
}
