package structmethodsinterface

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
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
