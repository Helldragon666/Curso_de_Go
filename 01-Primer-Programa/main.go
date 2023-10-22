package main

// main hace referencia a la carpeta actual
// package nombre_carpeta

// el archivo go.mod tiene las referencias de la version del lenguaje,
// asi como de librerias de terceros para cada proyecto

// para importar paquetes
//import "fmt"

// para importar multiples paquetes
import (
	"fmt"

	"github.com/Curso_de_Go/01-Primer-Programa/ejemplo"
)

// punto de entrada del programa
func main() {
	fmt.Println("Hola ;3")
	ejemplo.FuncionDeEjemplo()
}

/* Comandos:
1.- go run nombre_archivo -> correr el programa
2.- go build nombre_archivo -> compilar el programa
*/
