package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) DistToOrigin() float64 {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func main() {
	// create an object
	p1 := Point{3, 4}
	fmt.Println(p1.DistToOrigin())
}
