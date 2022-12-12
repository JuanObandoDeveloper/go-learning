package main

import "fmt"

func main() {
	o := 10
	p := 50

	// sum of two numbers
	sum := o + p
	fmt.Println(sum)

	// rest of two numbers
	rest := p - o
	fmt.Println(rest)

	// multiplication of two numbers
	multiplication := o * p
	fmt.Println(multiplication)

	// division of two numbers
	division := p / o
	fmt.Println(division)

	// module of two numbers
	module := p % o
	fmt.Println(module)

	// increment
	o++
	fmt.Println(o)

	// decrement
	p--
	fmt.Println(p)

	// Retos
	// 1. calcular area de rectangulo, trapecio y circulo
	var baseRect float64 = 12
	var alturaRect float64 = 18
	var areaRect float64 = baseRect * alturaRect
	fmt.Println(areaRect)

	var baseMayorTra float64 = 9.5
	var baseMenorTra float64 = 3.5
	var alturaTra float64 = 4
	areaTra := (baseMayorTra + baseMenorTra) / 2 * alturaTra
	fmt.Println(areaTra)

	var radioCir float64 = 5
	const pi float64 = 3.1416
	var areaCir float64 = pi * (radioCir * radioCir)
	fmt.Println(areaCir)

}
