package main

import (
	"fmt"
	"log"
	"strconv"
)

func isEven(a int) bool {
	return a%2 == 0
}

func login(user, pass string) bool {
	return user == "admin" && pass == "admin123"
}

func main() {
	// if conditional
	var v1 int = 1
	var v2 int = 2

	if v1 == 1 {
		fmt.Println("v1 is 1")
	} else {
		fmt.Println("v1 is not 1")
	}

	// AND (&&)
	if v1 == 1 && v2 == 2 {
		fmt.Println("it's true")
	}

	// OR (||)
	if v1 == 1 || v2 == 2 {
		fmt.Println("it's true")
	}

	// convert string to int
	value, err := strconv.Atoi("100")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)

	// challenge 1
	fmt.Println(isEven(2))
	fmt.Println(isEven(1))

	// challenge 2
	fmt.Println(login("admin", "admin123"))
	fmt.Println(login("admin", "admin"))

	// switch conditional
	switch module := 5 % 2; module {
	case 0:
		fmt.Println("it's even")
	default:
		fmt.Println("it's odd")
	}

	// no condition
	var v3 int = 200
	switch {
	case v3 > 100:
		fmt.Println("it's greater than 100")
	case v3 < 0:
		fmt.Println("it's less than 0")
	default:
		fmt.Println("it's between 0 and 100")
	}
}
