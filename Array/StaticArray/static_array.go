package main

import "fmt"

func main() {
	var numbers [5]int

	var _ []int // slice is dynamic array

	arr := [5]int{10, 20, 30, 40, 50}

	b := [...]int{1, 2, 3, 4, 5}

	numbers[0] = 11

	fmt.Println(numbers, arr, b)
}
