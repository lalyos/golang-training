package main

import "fmt"

type Point struct {
	X int
	Y int
}

func compair() {
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	p3 := Point{2, 1}
	fmt.Println(p1 == p2)
	fmt.Println(p1 == p3)
}

func main() {
	p1 := Point{1, 2}
	p2 := Point{X: 3, Y: 4}
	p3 := Point{X: 3}
	p4 := Point{}
	p5 := Point{
		X: 5,
		Y: 6,
	}
	p6 := &Point{3, 3}
	fmt.Println(p1, p2, p3, p4, p5, p6)
	//compair()
}
