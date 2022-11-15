package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"victor": 65,
		"vlad":   70,
		"angela": 63,
	}

	fmt.Println(ages)
	names := []string{}

	for k := range ages {
		names = append(names, k)
	}

	fmt.Println("names-b", names)
	sort.Strings(names)
	fmt.Println("names-a", names)

	for _, name := range names {
		fmt.Printf("%s \t %d \n", name, ages[name])
	}
}
