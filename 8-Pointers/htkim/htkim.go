package main

import "fmt"

func Swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	v1 := 1
	p1 := &v1
	fmt.Println(p1)  // 메모리 주소
	fmt.Println(v1)
	*p1 = 2  // 포인터에 값 할당
	fmt.Println(v1)
	p2 := new(int)  // 새 포인터 생성
	fmt.Println(p2)
	x := 10
	y := 20
	Swap(&x, &y)
	fmt.Println(x)
	fmt.Println(y)
}