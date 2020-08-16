/**
 * @Author: BookYao
 * @Description:
 * @File:  colorPoint
 * @Version: 1.0.0
 * @Date: 2020/8/16 23:29
 */

package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	x, y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(p.x - q.x, p.y - q.y)
}

func (p *Point) ScaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func main() {
	var cp ColoredPoint
	cp.x = 1
	cp.Point.y = 2
	fmt.Println(cp.Point.x, cp.y)

	red := color.RGBA{255, 0,0, 255}
	green := color.RGBA{0, 255, 0, 255}
	rp := ColoredPoint{Point{1, 1}, red}
	gp := ColoredPoint{Point{5, 4}, green}

	fmt.Println(rp.Distance(gp.Point))

	fmt.Println(rp.Point)
	rp.ScaleBy(2)
	gp.ScaleBy(2)
	fmt.Println(rp.Point)
	fmt.Println(rp.Distance(gp.Point))
}

  