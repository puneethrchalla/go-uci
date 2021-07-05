package main

import "fmt"

type MyInt int

// using receiver type to associate a method with a type (same package required)
func (mi MyInt) Double() int {
	return int(mi * 2)
}

func main() {
	// create the object (declaration)
	v := MyInt(3)
	// use dot notation to call the method
	fmt.Println(v.Double())
}
