package main

import (
	"fmt"
	"math"
)

func Sum(numbers ...int) int {
	c := make(chan int)
	go SumWithChannel(c, numbers...)
	return <-c
}

func SumWithChannel(c chan int, numbers ...int) {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	c <- sum
}

func DivideBy2(n int) (int, bool) {
	v := n / 2
	b := false
	if n % 2 == 0 {
		b = true
	}
	return v, b
}

func Max(numbers ...int) int {
	max := math.MinInt64
	for _, n := range numbers {
		if max < n {
			max = n
		}
	}
	return max
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fib(x uint) uint {
	if x < 2 {
		return x
	}
	return fib(x - 1) + fib(x - 2)
}

func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(DivideBy2(4))
	fmt.Println(Max(4, 2, 3, 41, 6))
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) // 1
	fmt.Println(nextOdd()) // 3
	fmt.Println(nextOdd()) // 5
	fmt.Println("=========================")
}