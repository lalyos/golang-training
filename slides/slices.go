package main

import "fmt"

func first3(a [5]int) []int {
	a[0] = 0
	return a[:3]
}

func first3p(a *[5]int) []int {
	a[0] = 0
	return a[:3]
}

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println(first3(a))
	fmt.Println(a)
}
