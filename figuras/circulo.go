package figuras

import (
	"fmt"
	"math"
)

type Circulo struct {
	punto Punto
	radio float32
}

func (c Circulo) Area() float32 {
	return c.radio * c.radio * math.Pi
}

func (c Circulo) Perimetro() float32 {
	return c.radio * 2 * math.Pi
}

func (c Circulo) ToString() string {
	cadena := fmt.Sprint("Circulo:", c)
	return cadena
}
