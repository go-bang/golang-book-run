package main

import (
	"testing"
)


type pair struct {
	in []int
	out []int
}

var testsSwap = []pair {
	{[]int{10, 20}, []int{20, 10}},
	{[]int{40, 10}, []int{10, 40}},
}

func TestSwap(t *testing.T) {
	for _, pair := range testsSwap {
		Swap(&pair.in[0], &pair.in[1])
		if pair.in[0] != pair.out[0] {
			t.Error(
				"For", pair.out[1],
				"expected", pair.out[0],
				"get", pair.in[0],
			)
		}
		if pair.in[1] != pair.out[1] {
			t.Error(
				"For", pair.out[0],
				"expected", pair.out[1],
				"get", pair.in[1],
			)
		}
	}
}