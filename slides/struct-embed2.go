package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 5
	w.Y = 5
	w.Radius = 10
	w.Spokes = 12
	fmt.Printf("%#v", w)
}
