package ejemplo

import "fmt"

func FuncionDeEjemplo() {
	fmt.Println("desde archivo_ejemplo.go ;33")
}

// NOTA: el scope de las variables y funciones va en relacion en si la primera letra de
// la variable / funcion es mayuscula o minuscula

// si la primera letra es mayuscula, la funcion/variable tendra un alcance global
// si la primera letra es minuscula, el alcance sera local
