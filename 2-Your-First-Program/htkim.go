package main

import (
	"fmt"
)

func f(i int) string {
	if i > 0 {
		return "a"
	} else {
		return "b"
	}
}

func main() {
	//fmt.Println("Hello World")
	//fmt.Printf("Hello, my name is %s", "htkim")
	//fmt.Println(f(-2))
	a := [...]int{1,2,3,4}
	b := []int{1,2,3,4}
	a1 := a
	b1 := b
	a[1] = 10
	b[1] = 10
	fmt.Println(a1)
	fmt.Println(b1)
}