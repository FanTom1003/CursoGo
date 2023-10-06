package objects

import "fmt"

// object car
type Car struct { // tipo con mayusculas es para acceso publico
	Marca string // propiedad con mayuscula es para acceso publico
	Year  int
}

func (car Car) String() string { // para personalizar el como se muestra un Struc
	return "este es un car"
}

// object car
type car struct { // tipo con minusculas es para acceso privado
	marca string // propiedad con minusculas es para acceso privado
	year  int
}

func Message() { // nombre con mayuscula es para acceso publico
	fmt.Println("mensaje")
}

func message() { // nombre con minusculas es para acceso privado
	fmt.Println("mensaje")
}
