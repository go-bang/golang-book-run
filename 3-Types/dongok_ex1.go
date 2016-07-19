package main

import "fmt"

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")

	fmt.Println("string length", len("aaaaaa"))

	//fmt.Println("sum=", sum)
	fmt.Println("input 계산")

	var a, b int
	var p string
	input_num, _ := fmt.Scanf("%d %s %d", &a, &p, &b)

	/*
	  , 는 %s 로 들어가서 하나의 입력으로 인지함.
	*/
	//_, _ = fmt.Scanf("%d,%s,%d", &a, &p, &b)

	fmt.Println("input number", input_num)
	fmt.Println("a=", a)
	fmt.Println("p=", p)
	fmt.Println("b=", b)

	ret := calc(a, p, b)
	fmt.Println("ret=", ret)
}

func calc(a int, p string, b int) (ret int) {
	switch p {
	case "+":
		ret = a + b
	case "-":
		ret = a - b
	case "/":
		ret = a / b
	case "*":
		ret = a * b

	default:
		ret = -999999999999
	}

	return ret
}
