package main

import (
	"fmt"
	"math"
)

const (
	rectType   = "RECT"
	circleType = "CIRCLE"
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

	//details(c)

	r := newGeometry(rectType, 2, 3)
	fmt.Printf("%+T\n", r)
	fmt.Println(r.area())
	fmt.Println(r.perim())

	c := newGeometry(circleType, 2)
	fmt.Printf("\n%+T\n", c)
	fmt.Println(c.area())
	fmt.Println(c.perim())

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

func newGeometry(geoType string, values ...float64) geometry {
	switch geoType {
	case rectType:
		return rect{width: values[0], height: values[1]}
	case circleType:
		return circle{radius: values[0]}
	}
	return nil
}
