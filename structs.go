package main

import (
	"fmt"

	pk "./mypackage"
)

type car struct {
	brand string
	model int
}

func main() {
	myCar := car{brand: "Ford", model: 2019}
	fmt.Println(myCar)

	// other way
	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.model = 2020
	fmt.Println(otherCar)

	// using mypackage public struct
	var anotherCar pk.CarPublic
	anotherCar.Brand = "Ferrari"
	anotherCar.Model = 2020
	fmt.Println(anotherCar)

	// using mypackage private struct
	// var otherCarPrivate pk.carPrivate // error cannot access private struct

	// using mypackage public function
	pk.Message("Hello from my package")
}
