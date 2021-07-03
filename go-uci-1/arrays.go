package main

import "fmt"

func main() {
	// Array Declaration
	var x [5]int
	// Array Literals
	var y [5]int = [5]int{1, 2, 3}
	z := [...]int{2, 4, 6} // length calculated from the number of initializers (values)

	fmt.Println(x)
	fmt.Println(y)

	for i, v := range z {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
}
