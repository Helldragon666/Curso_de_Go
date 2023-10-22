package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var err error

func IngresoDeDatos() {

	escaner := bufio.NewScanner(os.Stdin)

	fmt.Print("Ingrese un número: ")
	if escaner.Scan() {
		numero1, err = strconv.Atoi(escaner.Text())
		if err != nil {
			panic("el dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Print("Ingrese otro número: ")
	if escaner.Scan() {
		numero2, err = strconv.Atoi(escaner.Text())
		if err != nil {
			panic("el dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Println("Número 1 ingresado con valor:", numero1)
	fmt.Println("Número 2 ingresado con valor:", numero2)
}

func main() {
	IngresoDeDatos()
}
