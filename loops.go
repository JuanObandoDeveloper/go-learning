package main

import "fmt"

func main() {
	// for conditional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// for while
	var counter int = 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// for infinite
	var counter2 int = 0
	for {
		fmt.Println(counter2)
		counter2++
		if counter2 == 10 {
			break
		}
	}

	// for range
	var names = []string{"John", "Wick", "Doe"}
	for i, name := range names {
		fmt.Println(i, name)
	}

	// for range with _
	var names2 = []string{"John", "Wick", "Doe"}
	for _, name := range names2 {
		fmt.Println(name)
	}

	// for range with break
	var names3 = []string{"John", "Wick", "Doe"}
	for i, name := range names3 {
		fmt.Println(i, name)
		if i == 1 {
			break
		}
	}

	// for range with continue
	var names4 = []string{"John", "Wick", "Doe"}
	for i, name := range names4 {
		if i == 1 {
			continue
		}
		fmt.Println(i, name)
	}

	// challenge
	// 1. print 100 - 0
	for j := 100; j >= 0; j-- {
		fmt.Println(j)
	}

}
