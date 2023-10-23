package main

import (
	"fmt"
	"time"

	"github.com/Curso_de_Go/11-POO/estructuras"
	"github.com/Curso_de_Go/11-POO/herencia"
	in "github.com/Curso_de_Go/11-POO/interfaces" // alias en go
)

/*
Go es un lenguaje de programación que no tiene clases en el
sentido tradicional que se encuentran en lenguajes
orientados a objetos

En lugar de clases y herencia de clases, Go utiliza estructuras,
interfaces y composición (herencia) para lograr la encapsulación y la
reutilización de código.
*/

func main() {
	// Instancciar una estructura
	persona := estructuras.Persona{
		Nombre: "Yonatan",
		Edad:   22,
	}

	persona.Presentarse()

	persona2 := estructuras.Persona{"Ivan", 24}
	persona2.Presentarse()

	persona3 := new(estructuras.Persona)
	persona3.Nombre = "Juan"
	persona3.Edad = 30
	persona3.Presentarse()

	RegistrarUsuario()

	/////////////////////// Interfaces ///////////////////////////////
	rectangulo := in.Rectangulo{Base: 12, Altura: 30}
	ImprimirInformacion(rectangulo)

	/////////////////////// Herencia ///////////////////////////////
	estudiante := new(herencia.Estudiante)
	estudiante.Nombre = "Jacobo"
	estudiante.Edad = 23
	estudiante.Grado = 4

	estudiante2 := herencia.Estudiante{
		Humano: herencia.Humano{
			Nombre: "Antonio",
			Edad:   27,
		},
		Grado: 6,
	}

	fmt.Println(estudiante)
	fmt.Println(estudiante2)
}

func RegistrarUsuario() {
	usuario := new(estructuras.Usuario)
	usuario.CrearUsuario(1, "Yon", time.Now(), true)
	fmt.Println(*usuario)
}

// ///////////////////// Interfaces ///////////////////////////////
func ImprimirInformacion(c in.Calculos) {
	fmt.Printf("Área: %f\n", c.Area())
	fmt.Printf("Perímetro: %f\n", c.Perimetro())
}
