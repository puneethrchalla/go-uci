package main

import "fmt"

func main() {
	x := 6

	// Full if condition
	if x < 5 {
		fmt.Printf("x is less than %d", x)
	} else if x == 5 {
		fmt.Printf("x is equal to %d", x)
	} else {
		fmt.Printf("x is greater than %d", x)
	}

	// Switch-Case statement
	switch x {
	case 0:
		fmt.Println("x is 0")
	case 1:
		fmt.Println("x is 1")
	default:
		fmt.Println("neither 0 nor 1")
	}

}
