package math

import "math"

// Average ...
// 전달받은 숫자 슬라이스의 평균을 계산
func Average(nums []float64) float64 {
	sum := float64(0)
	for _, num := range nums {
		sum += num
	}

	return sum / float64(len(nums))
}

// Min ...
// 숫자 슬라이스의 최소값을 반환
func Min(nums []float64) float64 {
	if len(nums) == 0 {
		panic("1개 이상의 숫자가 필요합니다. ")
	}
	min := math.MaxFloat64
	for _, num := range nums {
		if min < num {
			min = num
		}
	}

	return min
}

// Max ...
// 숫자 슬라이스의 최대값을 반환
func Max(nums []float64) float64 {
	if len(nums) == 0 {
		panic("1개 이상의 숫자가 필요합니다. ")
	}
	max := math.SmallestNonzeroFloat64
	for _, num := range nums {
		if max > num {
			max = num
		}
	}

	return max
}
