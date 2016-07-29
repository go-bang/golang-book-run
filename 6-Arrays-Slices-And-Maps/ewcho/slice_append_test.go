package slice_test

import "fmt"

func Example_slice_append() {
	bar := []int{1, 2, 3, 4, 5}
	foo := append(bar, 6, 7) // Copy

	fmt.Println(bar)
	fmt.Println(foo)

	foo[3] = 10

	fmt.Println(bar)
	fmt.Println(foo)

	// Output:
	// [1 2 3 4 5]
	// [1 2 3 4 5 6 7]
	// [1 2 3 4 5]
	// [1 2 3 10 5 6 7]
}
