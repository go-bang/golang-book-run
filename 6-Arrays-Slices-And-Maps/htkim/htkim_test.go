package main

import (
	"testing"
)


type pair struct {
	in []int
	out int
}

var testsMin = []pair {
	{ []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}, 9},
}

func TestMin(t *testing.T) {
	for _, pair := range testsMin {
		v1 := Min(pair.in...)
		if v1 != pair.out {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"get", v1,
			)
		}
		v2 := MinFromArray(pair.in)
		if v2 != pair.out {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"get", v2,
			)
		}
	}
}