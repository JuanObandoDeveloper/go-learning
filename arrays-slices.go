package main

import (
	"fmt"
	"strings"
)

func isPalindrome(word string) bool {
	newWord := strings.ToLower(word)
	var reverse string
	for i := len(newWord) - 1; i >= 0; i-- {
		reverse += string(newWord[i])
	}
	return newWord == reverse
}

func main() {
	// arrays
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// slices
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))

	// slice methods
	fmt.Println(slice[0])
	fmt.Println(slice[1:2])
	fmt.Println(slice[:3])
	fmt.Println(slice[3:])

	// append
	slice = append(slice, 6)
	fmt.Println(slice)

	newSlice := []int{7, 8, 9}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// recorrido slice
	slice1 := []string{"hola", "que", "tal"}

	for i, valor := range slice1 {
		fmt.Println(i, valor)
	}

	// for _, valor := range slice1 {
	// 	fmt.Println(valor)
	// }

	// for i := range slice1 {
	// 	fmt.Println(i)
	// }

	// challenge
	fmt.Println(isPalindrome("ana"))
	fmt.Println(isPalindrome("hola"))
	fmt.Println(isPalindrome("Oso"))

}
