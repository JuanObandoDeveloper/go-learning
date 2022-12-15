package main

import "fmt"

func main() {
	// defer
	// execute after the function
	fmt.Println("Hola")
	defer fmt.Println("Mundo")

	// continue && break
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
		if i == 8 {
			break
		}
	}

}
