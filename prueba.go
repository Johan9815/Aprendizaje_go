package main

import "fmt"

type Persona struct {
	age  int
	name string
}

type Employee struct {
	id int
}

type FulltimeEmployee struct {
	Persona
	Employee
	endate int
}

type TemporaryEmployee struct {
	Persona
	Employee
	taxRate int
}

type PrintInfo interface {
	getMessage() string
}

func (f FulltimeEmployee) getMessage() string {

	return "Full Time employee"
}

func (t TemporaryEmployee) getMessage() string {

	return "Temporary Time employee"
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {

	//p := Persona{}
	//e := Employee{}

	full := FulltimeEmployee{}
	tem := TemporaryEmployee{}

	getMessage(full)
	getMessage(tem)

	//p.Setname("Johan")
	//fmt.Println(p.Getname())
	//e.Setid(1)
	//fmt.Println(e.Getid())
}

func (e *Employee) Getid() int {
	return e.id
}

func (e *Employee) Setid(id int) {
	e.id = id
}

func (e *Persona) Getname() string {
	return e.name
}

func (e *Persona) Setname(name string) {

	e.name = name
}
