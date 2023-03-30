package funciones

// Dada una cantidad cualquiera de valores enteros, devuelve cual de todos es el menor y cual el mayor
func MenorMayor(valores ...int)(int, int){
	minimo := valores[0]
	maximo := valores[0]

	for _, posicion := range valores[1:]{
		if posicion > maximo{
			maximo = posicion
		}
		if posicion < minimo{
			minimo = posicion
		}
	}

	return minimo, maximo
}