package main

import "fmt"

//structures

type Point struct {
	x        int32
	y        int32
	isOnGrid bool
}

type Circle struct {
	radius float64
	center *Point // we can also have it like *Point , omitting 'center'
}

func changeY(pt *Point) {
	pt.y = 100
}

func main() {
	var p1 Point = Point{1, 2, false}
	var p2 Point = Point{-5, 7, true}
	fmt.Println(p1.x, p2.x)
	fmt.Println(p1.y, p2.y)
	p3 := Point{5, 8, true}
	p4 := Point{x: 3}
	fmt.Println(p3, p4)
	p5 := &Point{y: -4}
	fmt.Println(p5)
	changeY(p5)
	fmt.Println(p5)

	p6 := &Point{y: 3}
	c1 := Circle{4.56, p6} //embedded struct
	fmt.Println(c1.center.x)

}
