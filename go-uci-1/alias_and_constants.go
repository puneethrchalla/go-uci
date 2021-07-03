package main

import "fmt"

func main() {
	// Declaring Multiple Alias in same block
	type (
		Celsius float32
		SSN     int32
	)

	// Declaring a single Alias
	type Weekday int

	// Declaring multiple Constants using iota (random values populated for each var)
	const (
		Monday Weekday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
	)

	// Declaring a single Constant
	const x = 1.3

	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Thursday)

}
