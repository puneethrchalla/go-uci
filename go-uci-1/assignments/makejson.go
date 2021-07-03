package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	user := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	// prompt for name and save it
	fmt.Println("Enter a name: ")
	scanner.Scan()
	user["name"] = scanner.Text()

	// prompt for address and save it
	fmt.Println("Enter an address: ")
	scanner.Scan()
	user["address"] = scanner.Text()

	// convert map to json
	json_data, err := json.Marshal(user)
	if err == nil {
		fmt.Println(string(json_data))
	}
}
