package main

import (
	"testing"
)

type pairShape struct {
	in  Shape
	out []float64
}

var testsShape = []pairShape{
	{&Circle{0, 0, 4}, []float64{50.26548245743669, 25.132741228718345}},
	{&Circle{0, 0, 9}, []float64{254.46900494077323, 56.548667764616276}},
	{&Rectangle{0, 0, 4, 6}, []float64{24, 20}},
	{&Rectangle{0, 0, 3, 3}, []float64{9, 12}},
}

func TestSwap(t *testing.T) {
	for _, pair := range testsShape {
		area_ := pair.in.area()
		if area_ != pair.out[0] {
			t.Error(
				"For", pair.in,
				"expected", pair.out[0],
				"get", area_,
			)
		}
		perimeter_ := pair.in.perimeter()
		if perimeter_ != pair.out[1] {
			t.Error(
				"For", pair.in,
				"expected", pair.out[1],
				"get", perimeter_,
			)
		}
	}
}