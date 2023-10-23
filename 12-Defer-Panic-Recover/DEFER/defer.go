package DEFER

import "fmt"

/*
utiliza para posponer la ejecución de una función hasta
que la función que la contiene termine
*/

func Funcion() {
	fmt.Println("Inicio de la función")
	defer fmt.Println("Este se imprimirá al final")
	fmt.Println("Fin de la función")
}
