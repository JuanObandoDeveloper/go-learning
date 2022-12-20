package main

import "fmt"

type shape interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	height float64
}

func (r rectangle) area() float64 {
	return r.base * r.height
}

func (s square) area() float64 {
	return s.base * s.base
}

func calc(s shape) float64 {
	return s.area()
}

func main() {
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 2, height: 4}

	fmt.Println(calc(mySquare))
	fmt.Println(calc(myRectangle))

	// lista de interfaces
	myInterface := []interface{}{"hola", 12, 4.90, true}
	fmt.Println(myInterface...)
}
