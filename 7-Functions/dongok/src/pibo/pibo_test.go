package pibo

import (
	"reflect"
	"testing"
)

func TestGetMemoriPiboList(t *testing.T) {
	t.Log("TestGetMemoriPiboList")
	a := GetMemoriPiboList(10)
	t.Log(a)
	correct := []uint{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	if reflect.DeepEqual(a, correct) {
		t.Error("GetMemoriPiboList")
	}
}
