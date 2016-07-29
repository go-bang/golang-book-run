package main

import "fmt"

func main() {
	sum := 10
	add := func(x, y int) int {
		sum += x + y
		return sum
	}

	fmt.Println(add(1, 2))
	fmt.Println(add(3, 4))
}
