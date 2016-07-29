package func_test

import "fmt"

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func ExampleDefer() {
	defer first() // 사용한 자원을 해제할 때 자주 사용
	second()

	// Output:
	// second
	// first
}
