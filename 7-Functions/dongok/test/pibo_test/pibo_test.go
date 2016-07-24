package pibo_test

import (
	"reflect"
	"testing"
	pibo "pibo"
)

func TestGetMemoriPiboList(t *testing.T) {
	t.Log("TestGetMemoriPiboList")
	a := pibo.GetMemoriPiboList(10)
	t.Log(a)
	correct := []uint{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	if !reflect.DeepEqual(a, correct) {
		t.Error("GetMemoriPiboList")
	}
}
