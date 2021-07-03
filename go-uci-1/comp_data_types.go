package main

import "fmt"

func main() {
	// Arrays
	x := [...]int{1, 2, 3, 4, 5}
	var y [5]int
	y[0] = 5
	y[1] = 20
	for i, v := range y {
		fmt.Printf("index: %d | value: %d\n", i, v)
	}
	x[4] = 10
	for i, v := range x {
		fmt.Printf("index: %d | value: %d\n", i, v)
	}

}
