package main

import (
	"fmt"
	"math/rand"
	"time"
)

func end(n int) {
	fmt.Println("end", n)
}

func f(n int) {
	defer end(n)
	for i := 0; i < 10; i++ {
		fmt.Println("thread", n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
