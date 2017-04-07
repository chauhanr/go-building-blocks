package interfaces

import "math"

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
