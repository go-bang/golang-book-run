package main

import "fmt"

func main() {
	// int, float
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)

	// string
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1]) // 101 문자가 바이트로 표현되기 때문
	fmt.Println("Hello" + "World")

	// boolean
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
