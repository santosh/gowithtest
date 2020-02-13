// Package shapes deal with geometrical calculations.
package shapes

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

// Circle represent a circle with a Radius.
type Circle struct {
	Radius float64
}
