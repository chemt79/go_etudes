package shapes

import (
	"fmt"
)

// Shape class
type Shape struct {
	x, y int
}

// NewShape is constructor of Shape
func NewShape(x, y int) *Shape {
	return &Shape{x, y}
}

// X is getter for shape.x
func (s *Shape) X() int {
	return s.x
}

// Y is getter for shape.y
func (s *Shape) Y() int {
	return s.y
}

// SetX is setter for shape.x
func (s *Shape) SetX(x int) {
	s.x = x
}

// SetY is setter for shape.y
func (s *Shape) SetY(y int) {
	s.y = y
}

// String is string view of Shape
func (s *Shape) String() string {
	return fmt.Sprintf("Shape x:%d, y:%d", s.x, s.y)
}
