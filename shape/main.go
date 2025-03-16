package main

import (
	"fmt"
	"math"
)

func main() {
	circle := Circle{radius: 5}
	rect := Rectangle{height: 10, width: 20}

	shapes := []Shape{rect, circle}

	for _, shape := range shapes {
		fmt.Printf("%.2f\n", shape.Area())
	}
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
