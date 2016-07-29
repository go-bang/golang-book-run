package main

import "fmt"

func ExamplePracticeMap() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}

	fmt.Println(x[2:5])

	y := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var min int = 100
	for _, value := range y {
		if min > value {
			min = value
		}
	}
	fmt.Println(min)

	// Output:
	// [c d e]
	// 9
}
