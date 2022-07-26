package Interface

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}
type Square struct {
	Side float64
}

type Triangle struct {
	Side1  float64
	Side2  float64
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (s Square) Area() float64 {
	return s.Side * s.Side
}
func (tri Triangle) Area() float64 {
	return (tri.Base * tri.Height) / 2
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
func (tri Triangle) Perimeter() float64 {
	return tri.Base + tri.Side1 + tri.Side2
}
func (s Square) Perimeter() float64 {
	return s.Side * s.Side
}
