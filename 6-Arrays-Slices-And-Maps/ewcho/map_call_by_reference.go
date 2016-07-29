package main

import "fmt"

func main() {
	bar := map[string]int{
		"a": 1,
		"b": 2,
	}

	foo := func(cha map[string]int) {
		cha["a"] = 3
	}

	foo(bar)

	fmt.Println(bar)
}
