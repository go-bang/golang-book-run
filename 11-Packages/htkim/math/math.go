package math

import "math"

//import "math"

// 인자로 전달된 숫자의 평균을 구한다
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// 인자로 전달된 숫자의 최소값을 구한다
func Min(xs []float64) float64 {
	min := math.MaxFloat64
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}

// 인자로 전달된 숫자의 최값을 구한다
func Max(xs []float64) float64 {
	max := -math.MaxFloat64
	for _, x := range xs {
		if max < x {
			max = x
		}
	}
	return max
}