package main

import "fmt"

/**
 * 정수를 받아 이등분해서 짝수면 true
 */
func half(i int) (int, bool) {
	if i%2 == 0 {
		return 1, true
	} else {
		return 0, false
	}
}

/**
 * 가변 매개변수를 받아서 가장 높은 수를 출력
 */
func findBiggestNumger(numList ...int) int {
	maxNum := 0
	for _, num := range numList {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

/**
 * 홀수를 반환하는 함수
 */
func makeOddGenerator() func() uint {
	var start uint
	start = 1

	return func() (ret uint) {
		ret = start
		start += 2
		return
	}
}

func fibonacci(n int) int {
	if n > 1 {
		return fibonacci(n-1) + fibonacci(n-2)
	} else {
		return n
	}
}

func ExampleMain() {
	fmt.Println(half(1000))
	fmt.Println(half(1001))

	fmt.Println(findBiggestNumger(10, 4, 2, 7, 15, 5, 8, 13, 11))

	oddNext := makeOddGenerator()
	fmt.Println(oddNext())
	fmt.Println(oddNext())
	fmt.Println(oddNext())

	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(50))

	// Output:
	// 1 true
	// 0 false
	// 15
	// 1
	// 3
	// 5
	// 1
	// 6765
}
