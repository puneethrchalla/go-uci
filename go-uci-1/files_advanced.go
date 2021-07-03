package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	barr := make([]byte, 20)
	// returns the number of bytes (nb) and moves the HEAD accordingly
	nb, err := f.Read(barr)
	if err == nil {
		fmt.Println(nb)
		fmt.Println(string(barr))
	}
}
