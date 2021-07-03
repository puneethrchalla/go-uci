package main

import "fmt"

const listLen = 10

func Swap(numlist []int, indx int) {
	temp := numlist[indx]
	numlist[indx] = numlist[indx+1]
	numlist[indx+1] = temp
}

func BubbleSort(numlist []int) {
	flag := 0
	for j := 0; j < len(numlist)-1; j++ {
		for i := 0; i < len(numlist)-j-1; i++ {
			if numlist[i] > numlist[i+1] {
				Swap(numlist, i)
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
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

	BubbleSort(numlist)
	fmt.Println(numlist)
}
