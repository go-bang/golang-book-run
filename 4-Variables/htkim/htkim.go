package main

import "fmt"

func Ftoc(f int) int {
	return (f - 32) * 5/9
}

func Fttom(ft float64) float64 {
	return ft * 0.3048
}

func main() {
	var v1 string = "first variable"
	v2 := "second variable"
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(Ftoc(86))  // 30
	fmt.Println(Fttom(25))  // 7.62
}