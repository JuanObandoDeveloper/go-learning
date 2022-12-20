package main

import (
	"fmt"
	"sync"
)

func speak(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	// better way
	var wg sync.WaitGroup

	// speak("Hola")
	fmt.Println("Hola")
	wg.Add(1)

	go speak("Mundo", &wg)

	wg.Wait()

	// not recomended, not efficient
	// time.Sleep(time.Second * 1)

	// example with goroutines
	go func(text string) {
		fmt.Println(text)
	}("Adios")
}
