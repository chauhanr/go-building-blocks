package interfaces

import (
	"math"
)

/** This file contains interface related to polygons and the methdos like area
  The interfaces have been created to check the instance type of the interface.
  test cases will check the interface type
*/

// Shaper is an interface that will be implemented by all the polygon structs here.
type Shaper interface {
	Area() float32
}

// Square is a ploygon with all sides equal and has 90 degree between sides.
type Square struct {
	side float32
}

// Circle is a ploygon with no side but a circumfrence around it.
type Circle struct {
	radius float32
}

// Area method that will make circle implement that Shaper interface
func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

//Area method that will make circle implement that Shaper interface
func (s *Square) Area() float32 {
	return s.side * s.side
}

// Classifier function will return the type of the interface.
func Classifier(items ...interface{}) []string {
	length := len(items)
	//log.Printf("%v\n", items)
	var interTypes = make([]string, length)
	for i, x := range items {
		switch x.(type) {
		case bool:
			interTypes[i] = "bool"
		case float32, float64:
			//log.Printf("setting the float item at index: %d \n", i)
			interTypes[i] = "float"
		case int, int64:
			//log.Printf("setting the int item at index: %d \n", i)
			interTypes[i] = "int"
		case nil:
			//log.Printf("setting the nil item at index: %d \n", i)
			interTypes[i] = "nil"
		case string:
			//log.Printf("setting the string item at index: %d \n", i)
			interTypes[i] = "string"
		default:
			//log.Printf("setting the undefined item at index: %d \n", i)
			interTypes[i] = "undefined"
		}
	}
	return interTypes
}
