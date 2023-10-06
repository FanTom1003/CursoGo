package src

import "fmt"

func varialbesConstantes() {
	//Declaración de constantes
	const pi float64 = 3.14 //con tipo
	const nombre = "Carlos" //sin tipo

	fmt.Println(pi, nombre)

	//Declaración de constantes
	base := 14          //si no se ha declaro, se declra y asigna un valor
	var altura int = 14 // se asigna el tipo de dato y valor
	var area int        // el tipo de dato

	fmt.Println(base, altura, area)

	//zero values
	var number int      //valor por defecto 0
	var decimal float64 //valor por defecto 0
	var text string     //valor por defecto ""
	var boolean bool    //valor por defecto false

	fmt.Println(number, decimal, text, boolean)
}
