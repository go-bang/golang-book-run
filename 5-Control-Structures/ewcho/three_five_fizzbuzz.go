package main

import "fmt"

func main() {
	var newLine bool

	for i := 1; i <= 100; i++ {
		newLine = false
		if i%3 == 0 {
			newLine = true
			fmt.Print("Fizz")
		}

		if i%5 == 0 {
			newLine = true
			fmt.Print("Buzz")
		}

		if newLine == true {
			fmt.Println()
		}
	}
}
