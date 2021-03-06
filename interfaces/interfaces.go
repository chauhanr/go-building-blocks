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

type specialString string

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
			interTypes[i] = "float"
		case int, int64:
			interTypes[i] = "int"
		case nil:
			interTypes[i] = "nil"
		case string:
			interTypes[i] = "string"
		case *Square:
			interTypes[i] = "Square"
		case specialString:
			interTypes[i] = "string"
		default:
			interTypes[i] = "undefined"
		}
	}
	return interTypes
}
