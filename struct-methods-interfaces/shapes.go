// Package shapes deal with geometrical calculations.
package shapes

import "math"

// Perimeter takes float64 width and heigth and returns perimeter of a rectangle.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area takes float64 width and heigth and returns calculated
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

// Rectangle is a quadrilateral with same length of opposite sides.
// By this defition a square is also a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area implements area for a Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represent a circle with a Radius.
type Circle struct {
	Radius float64
}

// Area implements area for a Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle represents a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area implements area for a Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// Shape interface acts as a middleware. Adding new geometry (shape) is as easy as
// implementing Area() of the shape.
type Shape interface {
	Area() float64
}
