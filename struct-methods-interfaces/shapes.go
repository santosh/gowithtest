// Package shapes deal with geometrical calculations.
package shapes

// Perimeter takes float64 width and heigth and returns perimeter of a rectangle.
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

// Area takes float64 width and heigth and returns calculated
func Area(width, height float64) float64 {
	return width * height
}
