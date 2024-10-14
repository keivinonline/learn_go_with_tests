package main

import "math"

type Shape interface {
	Area() float64
}
type Triangle struct {
	Width  float64
	Height float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle *Rectangle) (result float64) {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() (result float64) {
	return r.Width * r.Height
}

func (t Triangle) Area() (result float64) {
	return t.Width * t.Height / 2
}
func (c Circle) Area() (result float64) {
	return math.Pi * c.Radius * c.Radius
}
