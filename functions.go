package main

import "fmt"

// manera de declarar funciones en go:
// func <nombre_funcion>(<parametros>) <tipo_retorno> {
// 	<codigo>
// }

// funcion normal sin parametros
func normalFunction1() {
	fmt.Println("Hola mundo")
}

// funcion normal con un parametro
func normalFunction(message string) {
	fmt.Println(message)
}

// funcion con varios parametros
func triArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

// funcion que retorna algo
func returnValue(a int) int {
	return a * 2
}

// funcion que retorna mas de un valor
func returnMoreValues(a int) (c, d int) {
	return a, a * 2
}

// funcion principal
func main() {
	normalFunction1()
	normalFunction("Hola mundo")
	triArgument(1, 2, "Hi")
	var value int = returnValue(20)
	fmt.Println("value: ", value)
	var value1, value2 int = returnMoreValues(20)
	fmt.Println("value1: ", value1, "value2: ", value2)
}
