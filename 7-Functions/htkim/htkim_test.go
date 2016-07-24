package main

import (
	"testing"
)

type pairArrayInt struct {
	in  []int
	out int
}

type pairUintUint struct {
	in  uint
	out uint
}

type divideBy2Output struct {
	quotient int
	isEven   bool
}
type pairDivideBy2 struct {
	in  int
	out divideBy2Output
}

var testsSum = []pairArrayInt{
	{[]int{1, 2, 3, 4}, 10},
	{[]int{6, 7, 8}, 21},
}

var testsDivideBy2 = []pairDivideBy2{
	{1, divideBy2Output{0, false}},
	{2, divideBy2Output{1, true}},
}

var testsMax = []pairArrayInt{
	{[]int{1, 2, 3, 4}, 4},
	{[]int{6, 70, 71, 4, 9, 8}, 71},
}

var testsOddGenerator = []pairUintUint{
	{1, 1},
	{2, 3},
	{3, 5},
}

func TestSum(t *testing.T) {
	for _, pair := range testsSum {
		v := Sum(pair.in...)
		if v != pair.out {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"get", v,
			)
		}
	}
}

func TestDivideBy2(t *testing.T) {
	for _, pair := range testsDivideBy2 {
		quotient, isEven := DivideBy2(pair.in)
		if quotient != pair.out.quotient || isEven != pair.out.isEven {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"get", divideBy2Output{quotient, isEven},
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range testsMax {
		v := Max(pair.in...)
		if v != pair.out {
			t.Error(
				"For", pair.in,
				"expected", pair.out,
				"get", v,
			)
		}
	}
}

func TestOddGenerator(t *testing.T) {
	nextOdd := makeOddGenerator()

	for _, pair := range testsOddGenerator {
		v := nextOdd()
		if v != pair.out {
			t.Error(
				"For", pair.in, "th Odd",
				"expected", pair.out,
				"get", v,
			)
		}
	}
}