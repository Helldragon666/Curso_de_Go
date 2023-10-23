package herencia

/*
	Go utiliza la composición y la incorporación de tipos anónimos
	para lograr la reutilización de código
*/
type Humano struct {
	Nombre string
	Edad   int
}

// Definición de un tipo "Estudiante" que incorpora (compone) un tipo "Persona"
type Estudiante struct {
	Humano // Incorporación de la estructura "Humano"
	Grado  int
}
