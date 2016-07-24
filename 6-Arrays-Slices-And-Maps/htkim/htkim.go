package main

import (
	"fmt"
	"math"
)

func Min(numbers ...int) int {
	return MinFromArray(numbers)
}

func MinFromArray(numbers []int) int {
	min := math.MaxInt64
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
}

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println(Min(x...))
	fmt.Println(MinFromArray(x))
}