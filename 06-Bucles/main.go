package main

import "fmt"

// bucle infinito
//for {
//
//}

// nota: break y continue tambien se pueden utilizar

func main() {
	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}

	j := 11
	for j < 21 {
		fmt.Println(j)
		j++
	}
}
