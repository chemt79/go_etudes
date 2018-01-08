package shapes

import "fmt"

// Rectangle class
type Rectangle struct {
	Shape
	x1, y1 int
}

// NewRectangle is a constructor for Rectangle
func NewRectangle(x, y, x1, y1 int) *Rectangle {
	return &Rectangle{Shape{x, y}, x1, y1}
}

// SetX1 setter for x1
func (r *Rectangle) SetX1(x1 int) {
	r.x1 = x1
}

// SetY1 setter for y1
func (r *Rectangle) SetY1(y1 int) {
	r.y1 = y1
}

// X1 getter for x1
func (r *Rectangle) X1() (x1 int) {
	return r.x1
}

// Y1 getter for y1
func (r *Rectangle) Y1() (y1 int) {
	return r.y1
}

// String is a string view for Rectangle
func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle x:%d, y:%d, x1:%d, y1:%d", r.x, r.y, r.x1, r.y1)
}

// Parent return the superclass pointer
func (r *Rectangle) Parent() *Shape {
	return &r.Shape
}
