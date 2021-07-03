package main

import "fmt"

const listLen = 5

func Swap(numlist []int, indx int, minIndx int) {
	temp := numlist[indx]
	numlist[indx] = numlist[minIndx]
	numlist[minIndx] = temp
}

func SelectionSort(numlist []int) {
	for i := 0; i < len(numlist)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(numlist); j++ {
			if numlist[minIdx] > numlist[j] {
				minIdx = j
			}
		}
		Swap(numlist, i, minIdx)
	}
}

func main() {
	var numlist []int
	var temp int

	fmt.Printf("Enter %d integers:\n", listLen)

	for count := 0; count < listLen; count++ {
		fmt.Scanf("%d", &temp)
		numlist = append(numlist, temp)
	}

	SelectionSort(numlist)
	fmt.Println(numlist)
}
