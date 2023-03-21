package funciones

import "fmt"

// Ejercicio 1
func polinomio(coeficientes ...float32) []string {

	arreglo := make([]string, len(coeficientes)) // make() crea un nuevo slice de strings
	// con la cantidad de elementos igual a la cantidad de variables ingresadas

	for i := 0; i < len(coeficientes); i++ {

		if i > 1 {
			arreglo[i] = fmt.Sprintf(" %+.1f X**%v ", coeficientes[i], i) // (%+ .1f) Con este ejemplo el input sería "-3.1416" y el output sería "- 3.1 ""
		} else if i == 1 {
			arreglo[i] = fmt.Sprintf(" %+.1f X ", coeficientes[i])
		} else {
			arreglo[i] = fmt.Sprintf(" %.1f ", coeficientes[i])
		}
	}

	return arreglo

}