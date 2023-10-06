package src

import "fmt"

func saludo() {
	fmt.Println("hola")
}
func saludoParametro(texto string) {
	fmt.Println(texto)
}
func cuadrado(valor int) int {
	return valor * valor
}
func dobleReturn(a int) (int, int) {
	return a + a, cuadrado(a)
}
func operaciones(a, b int) (int, int, int, int) {
	return a + b, a * b, a - b, a / b
}
func funciones() {
	saludo()
	doble, cuadrado := dobleReturn(5)
	fmt.Println(doble, cuadrado)
	suma, multiplicacion, resta, division := operaciones(20, 5)
	fmt.Println(suma, multiplicacion, resta, division)

}
