package estructurasRepetitivas

import (
	"fmt"
)

func factorial(numero int) {
	factorial := 1

	for i := 1; i < int(numero); i++ {
		factorial += factorial * i
	}

	fmt.Printf("El factorial de %v es %v\n", numero, factorial)
}
