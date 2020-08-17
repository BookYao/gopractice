/**
 * @Author: BookYao
 * @Description: translate path
 * @File:  translateFunc
 * @Version: 1.0.0
 * @Date: 2020/8/17 10:10
 */

package main

import "fmt"

type Point struct {
	x, y float64
}

type Path []Point

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.x - q.x, p.y - q.y}
}

func (path Path) translateBy(offset Point, add bool) {
	var op func(q, p Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	path := Path{{1, 2}, {2, 4}}
	fmt.Println("path:", path)

	tmpPoint := Point{1, 2}
	path.translateBy(tmpPoint, true)
	fmt.Println("Translate Path:", path)

	path.translateBy(tmpPoint, false)
	fmt.Println("Translate Path:", path)
}

  