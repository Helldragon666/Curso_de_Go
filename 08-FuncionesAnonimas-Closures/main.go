package main

import "github.com/Curso_de_Go/08-FuncionesAnonimas-Closures/funciones"

func main() {
	funciones.Calculos()

	contador := funciones.Incrementador()
	println(contador())
	println(contador())
}
