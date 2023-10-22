package bloque

import (
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
