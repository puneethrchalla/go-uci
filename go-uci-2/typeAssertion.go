package main

import "fmt"

func DrawShapeWithSwitch(s Shape2D) {
	switch sh := s.(type) {
	case *Circle:
		fmt.Println("This is a Circle!")
		fmt.Println(sh.Perimeter())
	case Square:
		fmt.Println("This is a Square!")
		fmt.Println(sh.Perimeter())
	default:
		fmt.Println("Unknown Shape!")
	} 
}

func DrawShape(s Shape2D)  {
	sqr, ok := s.(Square)
	if ok {
		fmt.Println("This is a Square!")
		fmt.Println(sqr.Area())
	}
}
