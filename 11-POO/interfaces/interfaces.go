package interfaces

/*
 Interfaces -> permiten definir un conjunto de métodos que deben ser
 implementados por cualquier tipo que desee satisfacer la interfaz.
*/

type Calculos interface {
	Area() float64
	Perimetro() float64
}

///////////////////////////// Implementacion //////////////////////////
type Rectangulo struct {
	Base   float64
	Altura float64
}

func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}

func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.Base + r.Altura)
}
