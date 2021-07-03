package main

import "fmt"

func main() {
	var x int = 100
	var z int
	var y int
	var ptr *int // declaring an int pointer

	ptr = &x
	var ptr2 = new(int)
	y = *ptr + 100
	ptr2 = &z
	fmt.Println(ptr2)
	fmt.Println(ptr)
	fmt.Println(y)
	fmt.Println(&z)
}
