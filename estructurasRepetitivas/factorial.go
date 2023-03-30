package estructurasRepetitivas

// Factorial: dado un numero entero, devuelve su factorial
func Factorial(numero int) int{
	factorial := 1

	for i := 1; i < int(numero); i++ {
		factorial += factorial * i
	}
	
	return (factorial)
}
