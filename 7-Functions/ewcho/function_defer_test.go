package func_test

import "fmt"

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func ExampleDefer() {
	defer first()
	second()

	// Output:
	// second
	// first
}
