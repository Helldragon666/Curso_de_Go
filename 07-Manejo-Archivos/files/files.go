package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Curso_de_Go/07-Manejo-Archivos/funcion"
)

var rutaArchivo = "./files/txt/tabla.txt"

func GuardarEnArchivo() {
	texto := funcion.IngresarNumero()
	archivo, err := os.Create(rutaArchivo)

	if err != nil {
		fmt.Println("Error al crear el archivo", err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func ModificarArchivo() {
	texto := funcion.IngresarNumero()

	if !Unir(rutaArchivo, texto) {
		fmt.Println("Error al modificar el archivo")
	}
}

func Unir(ruta string, texto string) bool {
	archivo, err := os.OpenFile(ruta, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error al unir el contenido")
		return false
	}

	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Error al escribir el contenido")
		return false
	}

	archivo.Close()
	return true
}

func LeerArchivo() {
	//contenido, err := os.ReadFile(rutaArchivo)

	//if err != nil {
	//	fmt.Println("Error al leer el archivo", err.Error())
	//	return
	//}

	//fmt.Println(string(contenido))

	// otra forma
	archivo, err := os.Open(rutaArchivo)

	if err != nil {
		fmt.Println("Error al leer el archivo", err.Error())
		return
	}

	escaner := bufio.NewScanner(archivo)

	for escaner.Scan() {
		registro := escaner.Text()
		println(registro)
	}

	archivo.Close()
}
