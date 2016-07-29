package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 메모이제이션할 슬라이스 변수 확보
	// 피포나치 돌려서 끝
	memoization := make([]big.Int, 101)

	for i, _ := range memoization {
		if i <= 1 {
			memoization[i] = i
		} else {
			memoization[i] = memoization[i-1] + memoization[i-2]
		}
	}
	fmt.Println(memoization[92])
	fmt.Println(memoization[93]) // 93 부터 결과가 이상해짐(음수?)
	fmt.Println(memoization[94])
}
