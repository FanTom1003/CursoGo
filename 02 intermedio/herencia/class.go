package main

import "fmt"

type Person struct {
	id   int
	name string
}

type Employee struct {
	Person // herencia
	code   string
}

func main() {
	e := Employee{}
	e.id = 2
	e.name = "juan"
	e.code = "223"
	fmt.Println(e)
}
