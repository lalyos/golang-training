package main

import "fmt"

func main() {
	s1 := "1.line\n2.line\n3.line\r4.line"
	s2 := `1.line
2.line
3.line`

	fmt.Println(s1)
	fmt.Println("---")
	fmt.Println(s2)
}
