package estructuras

import (
	"fmt"
	"time"
)

type Persona struct {
	// atributos
	Nombre string
	Edad   int
}

// metodos
/*
	Estructura
	func (receptor Tipo) NombreDelMetodo() {
    // Código del método
	}

	1.- receptor es el nombre del parámetro que se utiliza para el tipo al que se asocia el método.
	2.- Tipo es el tipo de dato al que se asocia el método.

	NOTA: los metodos se declaran fuera de la estructura
*/

func (p Persona) Presentarse() {
	fmt.Printf("Hola, soy %s y tengo %d años.\n", p.Nombre, p.Edad)
}

/////////////////////////////////////////////////////////////////////////////

type Usuario struct {
	Id         int
	Nombre     string
	CreadoEn   time.Time
	Confirmado bool
}

func (u *Usuario) CrearUsuario(id int, nombre string, creadoEn time.Time, confirmado bool) {
	u.Id = id
	u.Nombre = nombre
	u.CreadoEn = creadoEn
	u.Confirmado = confirmado
}
