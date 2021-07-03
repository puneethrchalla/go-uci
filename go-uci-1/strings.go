package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var x int
	var y string
	// Conversion between data types
	x = 100
	y = strconv.Itoa(x)
	fmt.Println(y)
	// Test runes
	y = "Hello"
	fmt.Println(unicode.IsDigit(rune(y[0])))
	fmt.Println(strings.Contains(y, "Hell"))
	fmt.Printf(strings.ToLower(y) + "\n")
}
