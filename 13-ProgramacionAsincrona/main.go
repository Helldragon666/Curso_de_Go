package main

import "github.com/Curso_de_Go/13-ProgramacionAsincrona/gorutinas"

/*
La programación asincrónica en Go se puede lograr mediante gorutinas
(goroutines) y canales (channels).
*/

func main() {
	go gorutinas.SepararPalabras("Hola")
}
