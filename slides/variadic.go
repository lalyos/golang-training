package main

import "fmt"

func add(nums ...int) {
	fmt.Println("nums", nums)
	for i := range nums {
		fmt.Println("next:", nums[i])
		nums[i] = 0
	}

}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	add(s...)
	fmt.Println(s)
}
