package main

import "fmt"

func main() {
	// create the underlying array
	primes := [5]int{1, 2, 3, 4, 5}
	// declare & initialize a slice
	var s []int = primes[1:4]
	// functions on slice
	fmt.Println(primes)
	fmt.Println(s)
	fmt.Printf("Length of Slice: %d\n", len(s))
	fmt.Printf("Capacity of Slice: %d\n", cap(s))
	// Appending to a slice
	s = append(s, 9)
	fmt.Printf("Length of Slice: %d\n", len(s))
	fmt.Printf("Capacity of Slice: %d\n", cap(s))
	// Using make() - declare but not initialize
	//s1 := make([]int, 10)
	//s2 := make([]int, 10, 15)
}
