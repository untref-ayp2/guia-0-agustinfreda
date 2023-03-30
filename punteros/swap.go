package punteros

import "fmt"

func Swap(px, py *int) {
	fmt.Println("X:", *px)
	fmt.Println("Y:", *py)
	*px, *py = *py, *px
	fmt.Println("*-------------*")
}
