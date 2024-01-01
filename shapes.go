package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sidelenght float64
}

func runShapes() {

	t := triangle{
		base:   10,
		height: 10,
	}

	s := square{
		sidelenght: 10,
	}

	printArea(t)
	printArea(s)
}

func (s square) getArea() float64 {

	return s.sidelenght * s.sidelenght
}

func (t triangle) getArea() float64 {

	return 0.5 * t.base * t.height
}

func printArea(s shape) {

	fmt.Println(s.getArea())
}
