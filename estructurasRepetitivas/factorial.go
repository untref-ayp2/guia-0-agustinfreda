package estructurasRepetitivas


func Factorial(numero int) int{
	factorial := 1

	for i := 1; i < int(numero); i++ {
		factorial += factorial * i
	}
	
	return (factorial)
}
