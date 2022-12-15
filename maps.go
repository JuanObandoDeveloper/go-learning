package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Juan"] = 20
	m["Pedro"] = 30

	fmt.Println(m)

	// recorrer map
	for key, value := range m {
		fmt.Println(key, value)
	}

	// find one
	value, ok := m["Juan"]
	fmt.Println(value, ok)
}
