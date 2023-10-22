package retornos

import "strconv"

/*
Estructura:
	func nombre_funcion(parametro1 tipo_de_dato1, parametro2 tipo_de_dato2, ...) tipo_de_retorno {
		bloque de codigo
		return dato
	}

	Devolver multiples valores
		func nombre_funcion(parametro1 tipo_de_dato1, parametro2 tipo_de_dato2, ...) (tipo_de_retorno1, tipo_de_retorno2,...) {
		bloque de codigo
		return dato1, dato2, ... datoN
	}
*/

// va a devolver un bool y un string
func ConvertirATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
