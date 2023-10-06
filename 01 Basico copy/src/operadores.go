package src

import "fmt"

func operadores() {
	//Declaración de constantes
	var x = 36 //con tipo
	var y = 6  //sin tipo

	//suma
	result := x + y
	fmt.Println("suma", result)
	//resta
	result = x - y
	fmt.Println("resta", result)
	//multiplicacion
	result = x * y
	fmt.Println("multiplicacion", result)
	//division
	result = x / y
	fmt.Println("division", result)
	//modulo
	result = x % y
	fmt.Println("modulo", result)
	//incremental
	x++
	fmt.Println("incremental", result)
	//decremental
	x--
	fmt.Println("decremental", result)
}
