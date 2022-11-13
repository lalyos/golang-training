// START1 OMIT
package main

import (
	"fmt"
)

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
	*p++
	fmt.Println(*p)
}

// END1 OMIT

// START2 OMIT
func newInt1() *int {
	return new(int)
}

// END2 OMIT

// START3 OMIT
func newInt2() *int {
	var dummy int
	return &dummy
}

// END3 OMIT
