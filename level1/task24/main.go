package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (a *Point) Distance(b *Point) float64 {
	dist := (a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)
	return math.Sqrt(dist)
}

func main() {
	a := NewPoint(3, 4)
	b := NewPoint(0, 0)

	fmt.Println(a.Distance(b))
}
