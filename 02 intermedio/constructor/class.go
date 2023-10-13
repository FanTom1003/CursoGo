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
func newEmployee(id int, name string) *Employee {
	return &Employee{
		id:   id,
		name: name,
	}
}
func main() { // constructores de clases
	//1.
	e := Employee{}
	fmt.Println(e)
	//2
	e2 := Employee{
		id:   1,
		name: "frank",
	}
	fmt.Println(e2)
	//3
	e3 := new(Employee)
	fmt.Println(*e3)
	//4
	e4 := newEmployee(2, "juan")
	fmt.Println(*e4)
}
