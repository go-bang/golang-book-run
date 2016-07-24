package main

import (
	"testing"
)


type pairFtoc struct {
	in int
	out int
}

type pairFttom struct {
	in float64
	out float64
}

var testsFtoc = []pairFtoc {
	{ 86, 30},
}

var testsFttom = []pairFttom {
	{ 25.0, 7.62},
}

func TestFtoc(t *testing.T) {
	for _, pair := range testsFtoc {
		v := Ftoc(pair.in)
		if v != pair.out {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"get", v,
			)
		}
	}
}

func TestFttom(t *testing.T) {
	for _, pair := range testsFttom {
		v := Fttom(pair.in)
		if v != pair.out {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"get", v,
			)
		}
	}
}