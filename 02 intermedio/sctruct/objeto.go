package main

import "fmt"

type Employee struct { //object
	id   int    // attribute
	name string // attribute
}

func main() {
	e := Employee{}
	fmt.Println(e)
}
