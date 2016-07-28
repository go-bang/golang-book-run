package main

import (
	"fmt"
	"math/big"
	"time"
)

func makeFibonacciWithMemo() func(x int64) *big.Int {
	memo := make(map[int64]*big.Int)
	var fib func(x int64) *big.Int
	fib = func(x int64) *big.Int {
		n, exist := memo[x]
		if exist {
			return n
		} else {
			if x < 2 {
				memo[x] = big.NewInt(x)
			} else {
				memo[x] = new(big.Int)
				memo[x].Add(fib(x - 1), fib(x - 2))
			}
			return memo[x]
		}
	}
	return fib
}

func makeFibonacciWithMemoChannel() func(x int64, c chan *big.Int) {
	memo := make(map[int64]*big.Int)
	var fib func(x int64, c chan *big.Int)
	fib = func(x int64, c chan *big.Int) {
		n, exist := memo[x]
		if exist {
			c <- n
		} else {
			if x < 2 {
				memo[x] = big.NewInt(x)
			} else {
				memo[x] = new(big.Int)
				c1 := make(chan *big.Int)
				c2 := make(chan *big.Int)
				go fib(x - 1, c1)
				go fib(x - 2, c2)
				memo[x].Add(<-c1, <-c2)
			}
			c <- memo[x]
		}
	}
	return fib
}

func fibonacci(x int64) (*big.Int) {
	if x < 2 {
		return big.NewInt(x)
	}
	z := new(big.Int)
	return z.Add(fibonacci(x - 1), fibonacci(x - 2))
}

func fibonacciWithChannel(x int64, c chan *big.Int) {
	if x < 2 {
		c <- big.NewInt(x)
	} else {
		z := new(big.Int)
		c1 := make(chan *big.Int)
		//c2 := make(chan *big.Int)
		go fibonacciWithChannel(x - 1, c1)
		go fibonacciWithChannel(x - 2, c1)
		z.Add(<-c1, <-c1)
		c <- z
	}
}

func main() {
	limit := int64(30)
	printOn := true

	startFibonacci := time.Now()
	for i := int64(0); i <= limit; i++ {
		if printOn {
			fmt.Println(i, ": ", fibonacci(i))
		}
	}
	elapsedFibonacci := time.Since(startFibonacci)
	defer fmt.Println("fib: ", elapsedFibonacci)

	startFibonacciWithChannel := time.Now()
	cFibonacciWithChannel := make(chan *big.Int)
	for i := int64(0); i <= limit; i++ {
		go fibonacciWithChannel(i, cFibonacciWithChannel)
		if printOn {
			fmt.Println(i, ": ", <-cFibonacciWithChannel)
		}
	}
	elapsedFibonacciWithChannel := time.Since(startFibonacciWithChannel)
	defer fmt.Println("fib with chan: ", elapsedFibonacciWithChannel)

	startFibonacciWithMemo := time.Now()
	fibonacciWithMemo := makeFibonacciWithMemo()
	for i := int64(0); i <= limit; i++ {
		if printOn {
			fmt.Println(i, ": ", fibonacciWithMemo(i))
		}
	}
	elapsedFibonacciWithMemo := time.Since(startFibonacciWithMemo)
	defer fmt.Println("fib with memo: ", elapsedFibonacciWithMemo)

	startFibonacciWithMemoChannel := time.Now()
	fibonacciWithMemoChannel := makeFibonacciWithMemoChannel()
	cFibonacciWithMemoChannel := make(chan *big.Int)
	for i := int64(0); i <= limit; i++ {
		go fibonacciWithMemoChannel(i, cFibonacciWithMemoChannel)
		if printOn {
			fmt.Println(i, ": ", <-cFibonacciWithMemoChannel)
		}
	}
	elapsedFibonacciWithMemoChannel := time.Since(startFibonacciWithMemoChannel)
	defer fmt.Println("fib with memo&chan: ", elapsedFibonacciWithMemoChannel)
}



