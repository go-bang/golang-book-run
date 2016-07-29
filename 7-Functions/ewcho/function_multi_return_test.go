package function_test

import "fmt"

func multiReturn() (int, string) {
  return 123, "nice"
}

func multiNotSpecifyReturn() (ret int) {
  // 함수 시크니처를 명시할때 리턴할 변수를 선언할 수 있다
  ret = 10000
  return
}

func ExampleMultiReturn() {
	fmt.Println(multiReturn())
  fmt.Println(multiNotSpecifyReturn())

  // Output:
  // 123 nice
  // 10000
}
