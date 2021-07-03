package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func isInt(input string) bool {
	for _, char := range input {
		if !(unicode.IsDigit(rune(char))) {
			return false
		}
	}
	return true
}

func main() {
	var x string
	nums := make([]int, 0, 3)
	for {
		fmt.Println("Enter an integer:")
		fmt.Scan(&x)
		if x == "X" || x == "x" {
			break
		} else if isInt(x) {
			num, err := strconv.Atoi(x)
			if err == nil {
				nums = append(nums, num)
				sort.Ints(nums)
				fmt.Println(nums)
			}
		} else {
			fmt.Println("enter a valid integer or X to terminate")
		}
	}
}
