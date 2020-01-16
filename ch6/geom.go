package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p* Point) ScaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func main() {
	p1 := Point{x:0, y:0}
	p2 := Point{x:1,y:1}
	fmt.Println(p1.Distance(p2))

	p := Point{1, 2}
	(&p).ScaleBy(2)
	fmt.Println(p)

	(&Point{3,4}).ScaleBy(3)
}