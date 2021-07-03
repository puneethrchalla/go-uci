package main

import "fmt"

func main() {
	var floating float32
	fmt.Println("Enter a floating point ..")
	fmt.Scan(&floating)
	fmt.Println(int(floating))
}
