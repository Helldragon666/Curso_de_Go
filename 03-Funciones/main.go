package main

import (
	"fmt"

	"github.com/Curso_de_Go/03-Funciones/retornos"
)

func main() {
	status, texto := retornos.ConvertirATexto(123)
	fmt.Println(status)
	fmt.Println(texto)
}
