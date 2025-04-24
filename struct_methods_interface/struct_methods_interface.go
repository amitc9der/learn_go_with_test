package structmethodsinterface

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.width + rect.height)
}

func Area(rect Rectangle) float64 {
	return rect.width * rect.height
}
