package main

import (
	"fmt"
	"math"
)

// Calculate displacement using s = ½ a t2 + vot + so
func GenDisplaceFn(acceleration, initVelo, initDisp float64) func(float64) float64 {
	return func(time float64) float64 {
		return (0.5 * acceleration * math.Pow(time, 2)) + (initVelo * time) + initDisp
	}
}

func main() {
	inputFields := []string{"acceleration", "initial velocity", "initial displacement"}
	inputData := make(map[string]float64)
	var tempValue float64
	var time float64

	for _, input := range inputFields {
		fmt.Printf("Enter %s:\n", input)
		fmt.Scanf("%f", &tempValue)
		inputData[input] = tempValue
	}

	fn := GenDisplaceFn(inputData[inputFields[0]], inputData[inputFields[1]], inputData[inputFields[2]])
	fmt.Println("---------------------------------------------")
	fmt.Println("Values saved! Now, enter time!")
	fmt.Println("---------------------------------------------")
	for i := 0; i < 3; i++ {
		fmt.Println("Enter Time:")
		fmt.Scanf("%f", &time)
		fmt.Println("Displacement:")
		fmt.Println(fn(time))
	}
}
