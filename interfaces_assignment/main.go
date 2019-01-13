package main

import "fmt"

type area interface {
	printArea()
}

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base float64
}

type square struct {
	sidelength float64
}



func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) printArea() {
	fmt.Println(t.getArea())
}


func main() {
}