package main

import "fmt"

func main() {
	bar := "bar"
	if true {
		foo := "foo" // 블록단위 유효범위(scope)를 갖기 때문에 아래 Print 문에서 에러가 발생
	}

	for i := 0; i < 10; i++ {
		cha := "cha" // 블록단위 유효범위(scope)를 갖기 때문에 아래 Print 문에서 에러가 발생
	}

	fmt.Println(bar, foo, cha)
}
