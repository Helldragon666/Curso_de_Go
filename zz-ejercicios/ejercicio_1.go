package main

import (
	"fmt"
	"strconv"
)

func ConvertirAInt(valor string) (int, string) {
	//resultado, _ := strconv.Atoi(valor)

	resultado, err := strconv.Atoi(valor)

	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}

	if resultado > 100 {
		return resultado, "Es mayor a 100"
	} else {
		return resultado, "Es menor a 100"
	}
}

func main() {
	numero, texto := ConvertirAInt("90g")
	fmt.Println(numero)
	fmt.Println(texto)
}
