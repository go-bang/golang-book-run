package main

import "fmt"

func IsDivisible(n, m int) bool {
	return n % m == 0
}

func main() {
	for i := 1; i <= 100; i++ {
		if IsDivisible(i, 3) {
			fmt.Println(i)
		}
	}
	for i := 1; i <= 100; i++ {
		isMultiple3 := IsDivisible(i, 3)
		isMultiple5 := IsDivisible(i, 5)
		if isMultiple3 && isMultiple5 {
			fmt.Println("FizzBuzz")
		} else if isMultiple3 {
			fmt.Println("Buzz")
		} else if isMultiple5 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}