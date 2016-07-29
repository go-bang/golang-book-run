package map_test

import "fmt"

func ExampleMapDefine() {
	// var x map[string]int
	// panic: assignment to entry in nil map
	// Map 을 초기화 하지 않으면 panic 과 같은 Runtime exception 이 발생한다.

	x := make(map[string]int)

	x["a"] = 1
	x["b"] = 2

	fmt.Println(x)
	// Output:
	// map[a:1 b:2]
}

/*
 Map의 요소들은 순서가 없다.
 따라서 위 테스트 결과는 a 와 b가 바뀌어서 출력될 수 있다.

--- FAIL: ExampleMapDefine (0.00s)
got:
map[b:2 a:1]
want:
map[a:1 b:2]
FAIL
FAIL	command-line-arguments	0.007s
*/
