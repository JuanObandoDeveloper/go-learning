package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "pong")
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

func (myPC pc) String() string {
	return fmt.Sprintf("RAM: %d, Disk: %d, Brand: %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	var a int = 50
	var b *int = &a // b is a pointer to a

	fmt.Println(a) // 50
	fmt.Println(b) // memory address of a

	fmt.Println(b)  // memory address of a
	fmt.Println(*b) // value of a

	// modify value of a using pointer
	*b = 100
	fmt.Println(a) // 100

	// pointer to struct
	myPC := pc{ram: 16, disk: 500, brand: "Dell"}
	fmt.Println(myPC) // {16 500 Dell}

	myPC.ping() // Dell pong

	fmt.Println(myPC) // {16 500 Dell}

	myPC.duplicateRam()

	fmt.Println(myPC) // {32 500 Dell}

}
