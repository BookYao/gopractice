/**
 * @Author: BookYao
 * @Description: calculate point distance by function
 * @File:  pointDistance
 * @Version: 1.0.0
 * @Date: 2020/8/14 17:31
 */

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

type Path []Point

func distance(p, q Point) float64 {
	return math.Hypot(p.x - q.x, p.y - q.y)
}

func (p Point) distance(q Point) float64 {
	return math.Hypot(p.x - q.x, p.y - q.y)
}

func (p Path) distance() float64 {
	sum := 0.0
	for key := range p {
		if key > 0 {
			sum += p[key - 1].distance(p[key])
		}
	}
	return sum
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println("Normal Distance:", distance(p, q))
	fmt.Println("Funcion Distance:", p.distance(q))
	fmt.Println("Funcion Distance:", q.distance(p))

	point := Path {{1, 1}, {5, 1}, {5,4}, {1, 1}}
	fmt.Println("Function Path:", point.distance())
}

  