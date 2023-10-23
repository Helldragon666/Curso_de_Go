package funcion

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IngresarNumero() string {

	escaner := bufio.NewScanner(os.Stdin)
	var texto string

	fmt.Print("Ingrese un número: ")
	if escaner.Scan() {
		numero, err := strconv.Atoi(escaner.Text())

		for err != nil {
			IngresarNumero()
			break
		}

		for i := 1; i <= 10; i++ {
			texto += fmt.Sprintf("%d X %d = %d \n", numero, i, numero*i)
		}
	}

	return texto
}
