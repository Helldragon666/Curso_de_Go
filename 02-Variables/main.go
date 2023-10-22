package main

import "fmt"

/*
	Estructura:
		1.- var nombre_variable tipo_de_dato
		2.- var nombre_variable = valor
*/

func main() {
	var variable1 int
	var variable2 = 12
	variable3 := 3 // o tambien var variable3 int = 3

	fmt.Println("variable1:", variable1)
	fmt.Println("variable2:", variable2)
	fmt.Println("variable3:", variable3)
}
