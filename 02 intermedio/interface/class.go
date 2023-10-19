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

type FullEmployee struct {
	Employee
	date string
}

func (employee FullEmployee) showMessage() string {
	return "Full time employee"
}

type TempEmployee struct {
	Employee
	tax int
}

func (employee TempEmployee) showMessage() string {
	return "temp employee"
}

type PrintInfo interface { //interface
	showMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.showMessage())
}

func main() {
	e := Employee{}
	e.id = 2
	e.name = "juan"
	e.code = "223"
	fmt.Println(e)

	e2 := TempEmployee{}
	e3 := FullEmployee{}

	getMessage(e2)
	getMessage(e3)
}
