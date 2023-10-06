package main

import (
	"fmt"
	"time"
)

func main() {
	// repaso general
	/// declaracion de variables
	// var x int
	// x = 8
	// y := 8
	// fmt.Println(x)
	// fmt.Println(y)
	// // captura de errores
	// value, error := strconv.ParseUint("7", 0, 64)
	// if error != nil { //nil es la exprecion para null
	// 	fmt.Println(error)
	// } else {
	// 	fmt.Println(value)
	// }
	// // maps
	// mapa := make(map[string]int)

	// mapa["llave"] = 6
	// mapa["otro"] = 8

	// fmt.Println(mapa)
	// // slice
	// slice := []int{1, 2, 3}

	// for index, value := range slice {
	// 	fmt.Println(index)
	// 	fmt.Println(value)
	// }
	// slice = append(slice, 16)
	// for index, value := range slice {
	// 	fmt.Println(index)
	// 	fmt.Println(value)
	// }
	//go routines y apuntadores
	// las go rutines van de la mano con los canales para poder comunicar las rutinas
	canal := make(chan int)
	go doSomething(canal) //crea la go rutine
	<-canal
	//apuntadores
	g := 25
	fmt.Println(g)
	h := &g // apuntador de g
	fmt.Println(h)
	fmt.Println(*h) // * (operador estrella) muestra el valo almacenado en el apuntador
}

func doSomething(c chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("done")
	c <- 1
}
