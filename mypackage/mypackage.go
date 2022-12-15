package mypackage

import "fmt"

// CarPublic is a public struct
type CarPublic struct {
	Brand string
	Model int
}

type carPrivate struct {
	brand string
	model int
}

// Message is a public function
func Message(text string) {
	fmt.Println(text)
}
