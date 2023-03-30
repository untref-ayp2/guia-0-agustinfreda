package arreglos
/*
Escribir una función que reciba un arreglo de enteros como parámetros 
y devuelva la suma de todos sus elemento
*/
func SumaElementos(arreglo []int) int{
	suma := 0

	for i := 0; i < len(arreglo); i++{
		suma += arreglo[i]
	}

	return suma
}