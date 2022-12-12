package main

import "fmt"

func main() {
	// int numbers
	var a int   // it depends on the OS (32 or 64 bits)
	var b int8  // 8 bits (from -128 to 127)
	var c int16 // 16 bits (from -32768 to 32767)
	var d int32 // 32 bits (from -2147483648 to 2147483647)
	var e int64 // 64 bits (from -9223372036854775808 to 9223372036854775807)

	fmt.Println(a, b, c, d, e)

	// positive int numbers
	var f uint   // it depends on the OS (32 or 64 bits)
	var g uint8  // 8 bits (from 0 to 255)
	var h uint16 // 16 bits (from 0 to 65535)
	var i uint32 // 32 bits (from 0 to 4294967295)
	var j uint64 // 64 bits (from 0 to 18446744073709551615)

	fmt.Println(f, g, h, i, j)

	// float numbers
	var k float32 // 32 bits (from +- 1.18e-38 to +- 3.40e+38)
	var l float64 // 64 bits (from +- 2.23e-308 to +- 1.80e+308)

	fmt.Println(k, l)

	// string
	var m string // string of characters (UTF-8) using double quotes

	fmt.Println(m)

	// boolean
	var n bool // true or false

	fmt.Println(n)

	// complex numbers
	var o complex64  // float32 bits (real and imaginary parts)
	var p complex128 // float64 bits (real and imaginary parts)

	fmt.Println(o, p)

	q := 1 + 2i
	fmt.Println(q)

}
