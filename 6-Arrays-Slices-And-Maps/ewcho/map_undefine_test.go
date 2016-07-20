package map_test

import "fmt"

func ExampleMapUndefine() {
	bar := map[int]string{
		0: "Hello",
		1: "World",
	}

	fmt.Println(bar[2] == "")

	value, ok := bar[1]
	// 첫번째 반환: map 탐색 결과
	// 두번째 봔환: map 탐색 성공 여부

	fmt.Println(value, ok)
	// Output:
	// true
	// World true
}
