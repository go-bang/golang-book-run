package main

import (
	"fmt"
)

const ft float64 = 0.3048

func init() {
}
func main() {
	x := 5
	//x++ 가 더 효율적?
	x++
	fmt.Println(x)
	fmt.Println(ft)

	//섭씨를 화씨로
	f := 0
	var c float64
	c = float64(((f - 32) * 5 / 9))
	fmt.Println(c)

	//피트
	/*
		y := 10
		ftr := y / ft
		fmt.Println(ftr)
	*/

	var y float64 = 10
	ftr := y / ft
	fmt.Println(ftr)
}
