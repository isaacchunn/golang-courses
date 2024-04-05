package main

import "fmt"

//Assignment
/*

Write a program that creates two custom struct types called 'triangle' and 'square'
The 'square' type should be a struct with a field called 'sideLength' of type float64
The 'triangle' type should be a struct with a field called 'height' of type float64 and a field of type 'base' of type float64
Both types should have function called 'getArea' that returns the calculated area of the square or triangle
Area of a triangle = 0.5 * base * height.
Area of a square = sideLength * sideLength
Add a 'shape' interface that defines a function called 'printArea'. This function should calculate the area of the given shape and print it out to the terminal Design the interface so that the 'printArea' function can be called with either a triangle or a square.
*/

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	t := triangle{height: 1, base: 1}
	s := square{sideLength: 2}
	printArea(s)
	printArea(t)

}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}
