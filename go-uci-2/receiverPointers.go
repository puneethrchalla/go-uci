package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e Employee) SetId(id int) {
	e.id = id
	fmt.Println(e.id)
}

func (e Employee) DoubleId() {
	fmt.Println(e.id * 2)
}

func (e Employee) GetId() {
	fmt.Println(e.id)
}

func main() {
	v := Employee{}
	v.SetId(5)
	v.GetId()
	v.DoubleId()
}
