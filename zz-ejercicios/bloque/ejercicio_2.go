package bloque

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IngresarNumero() {
	escaner := bufio.NewScanner(os.Stdin)

	fmt.Print("Ingrese un número: ")
	if escaner.Scan() {
		numero, err := strconv.Atoi(escaner.Text())

		for err != nil {
			IngresarNumero()
			break
		}

		for i := 1; i <= 10; i++ {
			fmt.Println(numero, "X", i, "=", numero*i)
		}
	}
}
