package main

import (
	"fmt"
)

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
	getName(int) string
}

func main() {
	t1 := triangle{
		base:   5,
		height: 10,
	}
	t2 := triangle{
		base:   2.5,
		height: 5,
	}
	s1 := square{
		sideLength: 10,
	}
	s2 := square{
		sideLength: 5,
	}

	printArea(t1, 1)
	printArea(t2, 2)
	printArea(s1, 1)
	printArea(s2, 2)
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getName(num int) string {
	return "Triangle (" + fmt.Sprint(num) + ")"
}

func (s square) getName(num int) string {
	return "Square (" + fmt.Sprint(num) + ")"
}

func printArea(s shape, num int) {
	fmt.Println("Area of", s.getName(num), "is", s.getArea())
}
