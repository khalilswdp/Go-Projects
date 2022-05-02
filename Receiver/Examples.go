package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// InitMe Pointer to point because we want to modify the actual point not a copy of it (pass by copy)
func (p *Point) InitMe(xn, yn float64) {
	p.x = xn
	p.y = yn
}

// Print No pointer because we don't want to change the point (pass by pointer)
func (p Point) Print() {
	fmt.Printf("P.x: %f, P.y: %f\n", p.x, p.y)
}

func (p Point) DistToOrg() float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

func (p *Point) DoubleDistanceToOrg() {
	p.x = 2 * p.x
	p.y = 2 * p.y
}
func main() {
	var p Point
	p.InitMe(3, 4)

	p.Print()
	fmt.Println(p.DistToOrg())

	p.DoubleDistanceToOrg()

	p.Print()
	fmt.Println(p.DistToOrg())

}
