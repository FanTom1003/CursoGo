package interfaces

import "fmt"

type Figure2d interface {
	Area() float64
}

func CalcularArea(f Figure2d) {
	fmt.Println("Area:", f.Area())
}
