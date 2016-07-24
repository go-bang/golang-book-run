package main

import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
		fmt.Println(n, ":", i)
	}
}

func pinger(c chan <- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan <- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func Sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	//for i := 0; i < 10; i++ {
	//	go f(i)
	//}

	//c := make(chan string)
	//go pinger(c)
	//go ponger(c)
	//go printer(c)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			default:
				Sleep(time.Millisecond * 500)
				fmt.Println("nothing ready")
			}
			//fmt.Println(<-c1)
			//fmt.Println(<-c2)
		}
	}()
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}