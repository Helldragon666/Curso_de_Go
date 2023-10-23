package funciones

import "fmt"

func Calculos() {
	numero1 := 10
	numero2 := 30

	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}
	fmt.Printf("%d + %d = %d\n", numero1, numero2, suma(numero1, numero2))
}
