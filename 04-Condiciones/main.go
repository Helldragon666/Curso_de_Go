package main

import (
	"fmt"
	"runtime"
)

/*
	Estructura:
		if condicion {
			bloque de codigo verdadero
		} else {
			bloque de codigo falso
		}
*/

func main() {
	SistemaOperativo := runtime.GOOS
	if SistemaOperativo == "windows" {
		fmt.Println("Su sistema es windows")
	}

	valor := 1

	switch valor {
	case 1:
		fmt.Println("el valor es 1")
	case 2:
		fmt.Println("el valor es 2")
	}
}
