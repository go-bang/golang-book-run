package function_test

import "fmt"

func ExampleRaisePanic() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("Woops")
	// Output:
	// Woops
}
