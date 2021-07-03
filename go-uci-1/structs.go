package main

import "fmt"

func main() {
	type Person struct {
		name   string
		salary float32
		id     int32
	}

	var last_name string

	fmt.Println(last_name)

	p1 := new(Person)
	p2 := Person{name: "Puneeth Challa", salary: 112000.00}

	fmt.Println(p1.name)
	fmt.Println(p2.name)
}
