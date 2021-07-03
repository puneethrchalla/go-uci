package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//var userInput string
	fmt.Printf("Enter a String:\n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()

	userInput = strings.ToLower(userInput)
	userInput = strings.TrimSpace(userInput)

	if strings.HasPrefix(userInput, "i") && strings.HasSuffix(userInput, "n") && strings.Contains(userInput, "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
