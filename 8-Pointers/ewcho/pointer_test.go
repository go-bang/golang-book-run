package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
}

func swap(bar *int, foo *int) {
	*bar, *foo = *foo, *bar
}

func ExampleMain() {
	x := 1.5
	square(&x)
	fmt.Println(x)

	bar := 1
	foo := 2

	swap(&bar, &foo)

	fmt.Println(bar, foo)

	// Output:
	// 2.25
	// 2 1
}
