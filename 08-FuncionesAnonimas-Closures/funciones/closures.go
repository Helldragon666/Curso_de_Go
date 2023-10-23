package funciones

/*
Los closures son funciones anomimas que capturan
variables en su scope, en otras palabras convervan el
estado actual de esas variables
*/

// funcion que retorna una funcion que retorna un numero entero
func Incrementador() func() int {
	numero := 0

	return func() int {
		numero++
		return numero
	}
}
