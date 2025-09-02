package main

import (
	"fmt"
	"math"
)

// Shape is an interface that defines methods for calculating area and perimeter.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle represents a rectangle with width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of the rectangle.
func (this Rectangle) Area() float64 {
	return this.Width * this.Height
}

// Perimeter calculates the perimeter of the rectangle.
func (this Rectangle) Perimeter() float64 {
	return 2 * (this.Width + this.Height)
}

// Circle represents a circle with radius.
type Circle struct {
	Radius float64
}

// Area calculates the area of the circle.
func (this Circle) Area() float64 {
	return math.Pi * this.Radius * this.Radius
}

// Perimeter calculates the perimeter of the circle.
func (this Circle) Perimeter() float64 {
	return 2 * math.Pi * this.Radius
}

func main() {

	// create a rectangle instance
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Rectangle Area:", rect.Area())
	fmt.Println("Rectangle Perimeter:", rect.Perimeter())

	// create a circle instance
	c := Circle{Radius: 5}
	fmt.Println("Circle Area:", c.Area())
	fmt.Println("Circle Perimeter:", c.Perimeter())
}
