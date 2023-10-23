package main

import "fmt"

var miMapa map[string]int

var miMapa2 = map[string]int{
	"Yon":  22,
	"Ivan": 25,
	"Juan": 40,
}

var miMapa3 = make(map[string]int)

func MostrarMapas() {
	fmt.Println(miMapa)
	fmt.Println(miMapa2)

	// agregar elementos al map
	miMapa2["Carla"] = 34
	fmt.Println(miMapa2)

	// modificar el valor al map
	miMapa2["Carla"] = 10
	fmt.Println(miMapa2)

	// acceder a un valor del map
	fmt.Println(miMapa2["Ivan"])
	valor, existe := miMapa2["Carla"]
	fmt.Printf("%d existe? %t\n", valor, existe)

	// recorrer un map
	for nombre, edad := range miMapa2 {
		fmt.Println(nombre, edad)
	}

	//eliminar elementos (claves) del map
	delete(miMapa2, "Yon")
	fmt.Println(miMapa2)
}

func main() {
	MostrarMapas()
}
