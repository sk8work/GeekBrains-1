package interfaces

import "math"

// Square struct
type Square struct {
	Edge float64
}

// Area method from Square
func (s Square) Area() float64 {
	return s.Edge * s.Edge
}

// Perim method from Square
func (s Square) Perim() float64 {
	return 4 * s.Edge
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area method from Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perim method from Circle
func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.Radius
}

// Shape Interface operated methods of structs
type Shape interface {
	Area() float64
	Perim() float64
}

// SummAreas get summ of areas
func SummAreas(shapes ...Shape) float64 {
	res := 0.0
	for _, shape := range shapes {
		res += shape.Area()
	}
	return res
}

// SummPerims get summ of perims
func SummPerims(shapes ...Shape) float64 {
	res := 0.0
	for _, shape := range shapes {
		res += shape.Perim()
	}
	return res
}
