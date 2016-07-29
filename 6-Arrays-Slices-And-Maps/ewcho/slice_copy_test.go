package slice_test

import "fmt"

func Example_slice_copy() {
	bar := []int{1, 2, 3}
	foo := make([]int, 2)

	result := copy(foo, bar)
	fmt.Println(result)
	fmt.Println(foo)

	foo[1] = 10

	fmt.Println(bar)

	// Output:
	// 2
	// [1 2]
	// [1 2 3]
}
