package gorutinas

import (
	"fmt"
	"strings"
	"time"
)

/*
Las gorutinas son unidades de ejecución ligeras que se ejecutan
de manera concurrente y permiten que un programa realice múltiples
tareas de forma eficiente.
*/

func SepararPalabras(palabra string) {
	letras := strings.Split(palabra, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
