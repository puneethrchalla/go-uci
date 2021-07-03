package main

import "fmt"

func main() {

	// full for loop with iterator initialization, condition and update
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0

	// for loop used as while loop
	for j < 10 {
		fmt.Println(j)
		j++
	}
}
