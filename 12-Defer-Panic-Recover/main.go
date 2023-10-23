package main

import (
	"fmt"

	DEFER "github.com/Curso_de_Go/12-Defer-Panic-Recover/defer"
)

func main() {
	DEFER.Funcion()

	//resultado := PANIC.Dividir(10, 2)
	//fmt.Println("Resultado de la división:", resultado)

	//resultado = PANIC.Dividir(10, 0)                            // Esto provocará un panic
	//fmt.Println("Resultado de la segunda división:", resultado) // Esta línea nunca se ejecutará

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado de un pánico:", r)
		}
	}()

	fmt.Println("Antes del pánico")
	panic("Ocurrió un pánico")
	fmt.Println("Después del pánico") // Esta línea nunca se ejecutará
}
