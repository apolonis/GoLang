package main

import (
	"fmt"
	"math"
)

//Example interface's
type Geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry) {
	fmt.Println(g)

	fmt.Println(g.area())
	fmt.Println(g.perim())

}

func main() {

	r := Rect{width: 3, height: 4}

	c := Circle{radius: 5}

	measure(r)
	measure(c)

}
