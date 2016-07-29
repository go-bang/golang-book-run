package main

import (
	"fmt"
	"math"
)

func biggestBinaryNumber(digit int) string {
	var binary float64
	var result float64
	for i := 0; i < digit; i++ {
		binary += math.Pow(10, float64(i))
		result += math.Pow(2, float64(i))
	}

	return fmt.Sprintf("%d(%d)", int(binary), int(result))
}

func main() {
	// 가장 큰 8자리 숫자 (2진수)
	eightDigitBinary := biggestBinaryNumber(8)
	fmt.Println(eightDigitBinary)

	// 계산기
	fmt.Println(321325 * 424521)

	// 문자열 길이 구하기
	fmt.Println(len("안녕하세요."))

	fmt.Println(true && false)     // false
	fmt.Println(false && true)     // false
	fmt.Println(!(false && false)) // true
}
