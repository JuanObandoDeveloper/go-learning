package main

import "fmt"

func speak(text string, channel chan<- string) {
	channel <- text
}

func message(text string, channel chan<- string) {
	channel <- text
}

func main() {
	channel := make(chan string, 1)

	fmt.Println("Hola")

	go speak("Mundo", channel)

	fmt.Println(<-channel)

	// range, close and select
	channel2 := make(chan string, 2)
	channel2 <- "Hola"
	channel2 <- "Mundo"

	fmt.Println(len(channel2), cap(channel2))

	close(channel2) // close channel, no more data

	for message := range channel2 {
		fmt.Println(message)
	} // range channel, read all data

	// select
	channel3 := make(chan string)
	channel4 := make(chan string)
	go message("Hola", channel3)
	go message("Mundo", channel4)

	for i := 0; i < 2; i++ {
		select {
		case message1 := <-channel3:
			fmt.Println(message1)
		case message2 := <-channel4:
			fmt.Println(message2)
		default:
			fmt.Println("No message")
		}
	}
}
