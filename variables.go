package main

import "fmt"

func main() {
	// constants delcaration

	// untyped constants
	const pi = 3.14
	fmt.Println(pi)

	// typed constants
	const pi2 float64 = 3.1416
	fmt.Println(pi2)

	// multiple constants declaration
	const (
		a = 1
		b = 2
	)

	fmt.Println(a, b)

	// multiple constants declaration with type
	const (
		c int     = 3
		d float64 = 4.5
	)
	fmt.Println(c, d)

	// variables declaration

	// untyped variables
	var x = 5
	fmt.Println(x)

	// typed variables
	var y int = 6
	fmt.Println(y)

	// first declare then assign
	z := 7
	fmt.Println(z)

	// declare and asing
	var e int = 8
	fmt.Println(e)

	// declare but not assign
	var f int
	fmt.Println(f)

	// multiple variables declaration
	var (
		g = 9
		h = 10
	)
	fmt.Println(g, h)

	// multiple variables declaration with type
	var (
		i int     = 11
		j float64 = 12.5
	)
	fmt.Println(i, j)

	// zero value
	var k int
	var l float64
	var m bool
	var n string
	fmt.Println(k, l, m, n)

	// exercice
	const base = 10
	area := base * base
	fmt.Println(area)
}
