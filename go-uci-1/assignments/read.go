package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stringToByte(name string) []byte {
	field := make([]byte, 20)
	for indx, char := range []byte(name) {
		if indx >= 20 {
			break
		}
		field = append(field, char)
	}
	return field
}

func main() {
	// Get File Name
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the file name:")
	scanner.Scan()
	fileName := scanner.Text()

	// Declare the Struct & Slice
	type user struct {
		fname []byte
		lname []byte
	}
	var users []user

	// Read file into scanner obj
	fileReader, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File Read Failed!")
	}
	defer fileReader.Close()
	fileScanner := bufio.NewScanner(fileReader)

	for fileScanner.Scan() {
		name := strings.Split(fileScanner.Text(), " ")
		users = append(users, user{fname: stringToByte(name[0]), lname: stringToByte(name[1])})
	}

	fmt.Println("Full Names List:")

	for _, user := range users {
		fmt.Printf("%s %s\n", string(user.fname), string(user.lname))
	}

}
