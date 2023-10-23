package arreglos

import "fmt"

var miArreglo [10]int // arreglo con una longitud de 10

// otra forma de agregar valores
var numeros = [5]int{1, 2, 3, 4, 5}

var miMatriz [5][10]int

func AgregarValores() {
	miArreglo[4] = 10
	miArreglo[8] = 20

	fmt.Println(miArreglo)
}

func RecorrerArreglo() {
	for i := 0; i < len(miArreglo); i++ {
		fmt.Println(miArreglo[i])
	}
}

func MostrarMatriz() {
	fmt.Println(miMatriz)
}
