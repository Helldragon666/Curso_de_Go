package PANIC

/*
función panic se utiliza para indicar que algo inusual
y no manejable ha ocurrido en el programa.
*/

func Dividir(a, b int) int {
	if b == 0 {
		panic("Intento de división por cero")
	}
	return a / b
}
