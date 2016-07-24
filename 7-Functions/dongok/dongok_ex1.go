package main

import "fmt"

func main() {
	/*
	  1.sum은 숫자 슬라이스를 받아 그것들을 모두 합하는 함수다. Go에서 이 함수의 서명은 어떻게 될 것인가?

	  2.정수를 받아 그것을 이등분한 다음 짝수일 경우 참을 반환하고 홀수인 경우 거짓을 반환하는 함수를 작성하라. 예를 들어, half(1)은 (0, false)을 반환해야 하고 half(2)는 (1, true)를 반환해야 한다.
	  3.가변 매개변수를 하나 받아 숫자 목록에서 가장 큰 수를 찾는 함수를 작성하라.

	  4. makeEvenGenerator를 예제로 삼아 홀수를 생성하는 makeOddGenerator 함수를 작성하라.

	  5. 피보나치 수열은 fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2)로 정의된다. fib(n)를 구할 수 있는 재귀 함수를 작성하라.

	  5.지연, 패닉, 복구는 무엇인가? 런타임 패닉을 복구하는 방법은 무엇인가?
	*/

	sumSample := make([]int64, 5)
	sumSample[0] = 1
	sumSample[1] = 2
	sumSample[2] = 2
	sumSample[3] = 3
	sumSample[4] = 3

	fmt.Println(sum(sumSample))

	fmt.Println(half(11))

	fmt.Println(getMax(1, 2, 10, 4, 6, 7, 8))
	fmt.Println(getPiboList(30))
}

func getMax(a ...int) (max int) {
	/*
		sort.Ints(a)
		max = a[len(a)-1]
	*/
	for _, v := range a {
		if max < v {
			max = v
		}
	}
	return
}

func getPiboList(n int) []int {
	pibolist := make([]int, n)

	for i := 0; i < n; i++ {
		pibolist[i] = getPibo(i)
	}
	return pibolist
}

func getPibo(n int) int {
	x := 0
	if 0 <= n && n <= 1 {
		x = n
	} else {
		x = getPibo(n-1) + getPibo(n-2)
	}
	return x
}

func half(in int64) (h int64, isOdd bool) {
	h = in / 2
	isOdd = false
	if in%2 == 0 {
		isOdd = true
	}
	return
}

func sum(a []int64) (sumV int64) {
	for _, v := range a {
		sumV += v
	}
	return
}
