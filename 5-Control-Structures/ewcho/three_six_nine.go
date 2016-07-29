package main

import "fmt"

// 1과 100 사이의 숫자 중 3으로 나누어 떨어지는 수(3, 6, 9, ...)를 모두 출력하는 프로그램을 작성하라.
func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
