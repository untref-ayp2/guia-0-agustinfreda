package main

import (
	"fmt"

	"untref.agus.guia0/punteros"
)

func main() {
	x := 10
	y := 1
	var px = &x
	var py = &y
	punteros.Swap(px, py)
	fmt.Println("X:", x)
	fmt.Println("Y:", y)
}
