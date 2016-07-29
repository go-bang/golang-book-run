package main

import (
	"testing"
)


type pair struct {
	in []int
	out bool
}

var testsIsDivisible = []pair {
	{ []int{4, 2}, true},
	{ []int{9, 3}, true},
	{ []int{10, 3}, false},
}

func TestIsDivisible(t *testing.T) {
	for _, pair := range testsIsDivisible {
		v := IsDivisible(pair.in[0], pair.in[1])
		if v != pair.out {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"get", v,
			)
		}
	}
}