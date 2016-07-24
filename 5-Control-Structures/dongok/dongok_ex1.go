package main

import "fmt"

func main() {
	i := 10
	if i > 10 {
		fmt.Println("크다")
	} else {
		fmt.Println("작다")
	}

	/*
		for x := 1; x <= 100; x = x + 3 {
			fmt.Println(x)
		}
	*/
	for x := 1; x <= 100; x++ {
		fmt.Print("x", x, "\n")
		if x%3 == 0 {
			fmt.Print("Fizz")
		}
		if x%5 == 0 {
			fmt.Print("Buzz")
		}
		fmt.Print("\n")
	}
}
