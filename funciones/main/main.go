package main

import (
	"fmt"

	"untref.agus.guia0/funciones"
)

func main() {
	fmt.Println(funciones.Polinomio(2.0, -3.0, 1.5))
	fmt.Println(funciones.Polinomio(0, 2.0, -9.0, 3))
	fmt.Println(funciones.Polinomio(0, 0, 0))

	// funciones.Menu()
	
	fmt.Println(funciones.MenorMayor())
}
