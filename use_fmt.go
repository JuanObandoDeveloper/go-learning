package main

import "fmt"

func main() {
	// var declaration
	var helloMsg string = "Hello"
	var worldMsg string = "World"

	// println hello world
	fmt.Println(helloMsg, worldMsg)

	// printf hello world
	var nombre string = "Juan"
	var edad int = 20
	fmt.Printf("Hola %s tienes %d años\n", nombre, edad)
	fmt.Printf("Hola %v tienes %v años\n", nombre, edad)

	// sprintf hello world
	var message string = fmt.Sprintf("Hola %s tienes %d años", nombre, edad)
	fmt.Println(message)

	// data types
	fmt.Printf("helloMsg: %T\n", helloMsg)
	fmt.Printf("edad: %T\n", edad)

	// important

	// general:
	// %v -> value
	// %T -> type
	// %#v -> value with type
	// %% -> a literal % consumes no value

	//boolean:
	// %t -> boolean

	// integer:
	// %b -> binary
	// %d -> int decimal
	// %o -> octal
	// %O -> octal with 0o prefix
	// %x -> hexadecimal with lower-case letters fot a-f
	// %X -> hexadecimal with upper-case letters fot A-F
	// %U -> Unicode format: U+1234; same as "U+%04X"
	// %q -> a single-quoted character literal safely escaped with Go syntax.

	// floating-point and complex constituents:
	// %b -> decimalles scientific notation with exponent a power of two, in the manner o strconv.FormatFloat with the 'b' format, e.g. -123456p-78
	// %e -> scientific notation, e.g. -1.234456e+78
	// %E -> scientific notation, e.g. -1.234456E+78
	// %f -> decimal point but no exponent, e.g. 123.456
	// %F -> synonym for %f
	// % g -> %e for large exponents, %f otherwise
	// %G -> %E for large exponents, %F otherwise
	// %x -> hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
	// %X -> upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

	// string and slice of bytes (treated equivalently with these verbs):
	// %s -> the uninterpreted bytes of the string or slice
	// %q -> a double-quoted string safely escaped with Go syntax
	// %x -> base 16, lower-case, two characters per byte
	// %X -> base 16, upper-case, two characters per byte

	// slice:
	// %p -> base 16 notation, with leading 0x

	// pointer:
	// %p -> base 16 notation, with leading 0x

}
