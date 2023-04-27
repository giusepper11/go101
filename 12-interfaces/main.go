package main

import (
	"errors"
	"fmt"
	"math"
)

type geometric interface {
	area() float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(g geometric) {
	fmt.Println(g.area())
}

func ShowError(err error) {
	fmt.Println(err.Error())
}

func main() {
	fmt.Println("Starting...")

	rec := rectangle{
		width:  1,
		height: 2,
	}

	circ := circle{
		radius: 3,
	}

	getArea(rec)
	getArea(circ)

	ShowError(errors.New("error test"))
	
	// Empty interface
	var list []interface{}

	list = append(list, 10)
	list = append(list, true)
	list = append(list, "test")

	for _, v := range list {
		fmt.Println(v)
	}

}
