package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base, height float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{base: 4.0, height: 3.0}
	s := square{sideLength: 5.0}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}