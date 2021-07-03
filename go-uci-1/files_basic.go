package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat := []byte("My Name is Abhigna Manda")
	err := ioutil.WriteFile("test.txt", dat, 0777)
	if err == nil {
		fmt.Println("File Written!")
	}
}
