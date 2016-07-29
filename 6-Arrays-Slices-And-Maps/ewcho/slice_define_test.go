package slice_test

import "fmt"

func ExampleSliceDefine() {
	var x []string
	// panic: runtime error: index out of range
	// 길이가 0인 슬라이스가 생성된다. 따라서 아무 값도 넣을 수가 없다.

	x[0] = "Hello"
	x[1] = "World"

	fmt.Println(x)
	// Output:
	// [Hello World]
}

func ExampleSliceInitialize() {
	y := make([]string, 2)
	y[0] = "Hi"
	y[1] = "Bye"

	fmt.Println(y)
	// Output:
	// [Hi Bye]

}
