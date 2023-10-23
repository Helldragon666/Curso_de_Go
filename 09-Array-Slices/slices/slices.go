package slices

import "fmt"

/*
La diferencia entre un arreglo y un slice es que el arreglo tiene
un tamaño fijo y el slice es de longitud dinamica y flexible
*/

var miSlice []int
var nombres = []string{"Yonatan", "Juan", "Ivan"}
var numerosSlice = make([]int, 5)
var slice = make([]int, 5, 20) // slice con una longitud de 5, pero con una capacidad de 20

var arreglo = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func MotrarSlice() {
	miSlice = append(miSlice, 2) // agregar valores al final del slice
	miSlice = append(miSlice, 3)
	fmt.Println(miSlice)
}

func RecortarSlice() {
	porcion1 := arreglo[3:]
	porcion2 := arreglo[:5]
	porcion3 := arreglo[5:7]

	fmt.Println(porcion1)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func CapacidadDeSlice() {
	capacidad := cap(nombres) //capacidad de memoria que el lenguaje reserva para el slice
	fmt.Println(capacidad)

	capacidad2 := cap(slice)
	fmt.Println(capacidad2)
}

/*
NOTA: la capacidad de los slices va aumentando a medida que se añade un elemento,
esto con el fin de hacer los procesos mas optimos y mejorar el rendimiento
de las aplicaciones
*/
