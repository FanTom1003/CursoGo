package main

import "fmt"

type Employee struct { //object
	id   int    // attribute
	name string // attribute
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}
func (e *Employee) GetName() string {
	return e.name
}
func main() {
	e := Employee{}
	fmt.Println(e)
}
