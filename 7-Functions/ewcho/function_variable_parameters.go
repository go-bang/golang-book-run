package main

import "fmt"

func sample(args ...int) {
	for _, value := range args {
		fmt.Println(value)
	}
}

func main() {
	sample(1, 2, 3, 4, 5)
	sample([]int{10, 9, 8, 7, 6}...)
}
