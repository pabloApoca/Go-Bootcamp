package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

type geometry interface {
	area() float64
	perim() float64
}

func main() {

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	details(r)
	details(c)

}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func detailsC(c circle) {
	fmt.Println(c)
	fmt.Println(c.area())
	fmt.Println(c.perim())
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func details(g geometry) {
	fmt.Printf("%+T", g)

	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
