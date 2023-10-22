package datos

import (
	"fmt"
	"time" // para fechas
)

var CadenaGlobal string
var cadenaLocal string // no se podra usar en otros archivos

var BooleanoGlobal bool
var FlotanteGlobal float32
var FechaGlobal time.Time

func MostrarDatos() {
	CadenaGlobal = "mi texto ñ"
	BooleanoGlobal = true
	FlotanteGlobal = 23.67
	FechaGlobal = time.Now()
	fmt.Println(CadenaGlobal)
	fmt.Println(BooleanoGlobal)
	fmt.Println(FlotanteGlobal)
	fmt.Println(FechaGlobal)
}
