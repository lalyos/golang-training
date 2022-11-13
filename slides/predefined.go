package main

import "fmt"

// START OMIT
type int bool            // HL
var nil string = "nulla" // HL

func main() {
	var x int
	var str string = nil
	fmt.Println("x:", x)
	fmt.Println("str:", str)
}

//END OMIT
