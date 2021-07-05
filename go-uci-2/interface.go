package main

import (
	"fmt"
	"math"
)

// defining interface
type Shape2D interface {
	Area() float64
	Perimeter() float64
}

// defining concrete types
type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

// type "Square" satisfies the interface "Shape2D" (as it implements all it's methods
func (s Square) Area() float64 {
	return math.Pow(s.side, 2)
}
func (s Square) Perimeter() float64 {
	return s.side * 4
}

// type "Circle" satisfies the interface "Shape2D" (as it implements all it's methods
func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	// interface with Dynamic type & value
	var i Shape2D
	s1 := Circle{radius: 5}
	// as Circle satisfies Shape2D interface,
	i = &s1
	describe(i)
	// calling the method of dynamic type
	fmt.Println(i.Area())

	// Interface with nil Dynamic Value
	var ni Shape2D
	var s2 *Circle
	ni = s2
	describe(ni)

	// Nil interface value
	var ni2 Shape2D
	describe(ni2)

	// Empty Interface (can be used to pass any type)
	emptyInterface(5)
	emptyInterface(true)
	emptyInterface("hello")

	// Type Assertions
	DrawShape(Square{side: 4})
	DrawShapeWithSwitch(&Circle{radius: 7})
}

func describe(intf Shape2D) {
	fmt.Printf("(%v, %T)\n", intf, intf)
}

func emptyInterface(i interface{}) {
	fmt.Println(i)
}

// UseCase of Interface (Pool in Yard)
func FitInYard(poolShape Shape2D) bool {
	if poolShape.Area() > 100 && poolShape.Perimeter() > 100 {
		return true
	}
	return false
}
