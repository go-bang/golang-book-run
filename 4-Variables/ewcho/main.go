package main

import "fmt"

func transferTemperature(f float64) float64 {
	return (f - 32) * 5 / 9
}

func transferDistance(ft float64) float64 {
	return ft * 0.3048
}

func main() {
	var input string

	fmt.Println("무엇을 변환 하시겠습니까?")
	fmt.Println("1. 온도(T)")
	fmt.Println("2. 거리(D)")

	fmt.Scanf("%s", &input)

	if input == "1" || input == "T" {
		// 예제 프로그램을 출발점으로 삼아 화씨를 섭씨로 변환하는 프로그램을 작성하라. (C = (F - 32) * 5/9))
		var f float64

		fmt.Println("화씨 온도를 입력하세요.")
		fmt.Scanf("%f", &f)

		result := transferTemperature(f)

		fmt.Printf("%f℃", result)
		fmt.Println()
	} else if input == "2" || input == "D" {
		// 피트를 미터로 변환하는 프로그램을 하나 더 작성하라. (1 ft = 0.3048 m)
		var ft float64
		fmt.Println("피트 거리를 입력하세요.")
		fmt.Scanf("%f", &ft)

		result := transferDistance(ft)

		fmt.Printf("%fm", result)
		fmt.Println()
	} else {
		main()
	}
}
