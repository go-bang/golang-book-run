package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * (c.r * c.r)
}

func (c *Circle) perimeter() float64 {
	return math.Pi * (2 * c.r)
}

type Rectangle struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

func (re *Rectangle) distance(x1, y1, x2, y2 float64) float64 {
	x := x2 - x1
	y := y2 - y1
	return float64(math.Sqrt(x*x + y*y))
}

func (re *Rectangle) area() float64 {
	width := re.distance(re.x1, re.y1, re.x2, re.y1)
	height := re.distance(re.x1, re.y1, re.x1, re.y2)
	return width * height
}

func (re *Rectangle) perimeter() float64 {
	width := re.distance(re.x1, re.y1, re.x2, re.y1)
	height := re.distance(re.x1, re.y1, re.x1, re.y2)

	return 2*width + 2*height
}

type Shape interface {
	area() float64
	perimeter() float64
}

func main() {
	c := Circle{5}
	r := Rectangle{x1: 0, y1: 0, x2: 2, y2: 3}

	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())

}
