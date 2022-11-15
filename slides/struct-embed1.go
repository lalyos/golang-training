package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Center.X = 5
	w.Circle.Center.Y = 5
	w.Circle.Radius = 10
	w.Spokes = 12
	fmt.Printf("%#v", w)
}
