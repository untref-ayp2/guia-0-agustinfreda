package estructurasRepetitivas



func EsPrimo(num int) bool{
	primo := false

	if num <= 1{
		primo = false
		return primo
	}

	for i := 2; i < num; i++{
		if num % i == 0{
			primo = false
			return primo
		}
		primo = true
	} 

	return primo
}