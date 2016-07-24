package pibo

import (
	"fmt"
	"reflect"
)

func uintSort(unsorted []uint) []uint {
	var tmp uint

	fmt.Println("notsorted", unsorted, reflect.TypeOf(unsorted))
	unsorted2 := unsorted

	for i := 0; i < (len(unsorted2) - 1); i++ {

		for j := i + 1; j < len(unsorted2); j++ {

			if unsorted2[i] > unsorted2[j] {

				tmp = unsorted2[j]
				unsorted2[j] = unsorted2[i]
				unsorted2[i] = tmp

			}
		}
	}

	fmt.Println("uintSort", unsorted)
	fmt.Println("uintSort2", unsorted2)
	return unsorted
}

func GetRecursivePiboList(n uint) []uint {
	var pibolist []uint

	var i uint
	for i = 0; i <= n; i++ {
		pibolist = append(pibolist, getPibo(i))
	}
	return pibolist
}

func GetMemoriPiboList(n uint) []uint {
	var pibolist []uint
	retList := make(map[uint]uint)
	var v uint
	var i uint
	for i = 0; i <= n; i++ {
		v = getPiboWithMemo(i, retList)
		pibolist = append(pibolist, v)
		retList[i] = v
	}

	var keys []uint
	for k := range retList {
		keys = append(keys, k)
	}

	fmt.Println("a", keys)
	uintSort(keys)
	fmt.Println("b", keys)

	for k := range keys {
		v := retList[uint(k)]
		fmt.Println("key", k, "value", v)
	}

	return pibolist
}

func getPiboWithMemo(n uint, retList map[uint]uint) uint {
	value, ok := retList[n]
	if ok {
		return value
	}

	var x uint
	if 0 <= n && n <= 1 {
		x = n
	} else {
		x = getPiboWithMemo(n-1, retList) + getPiboWithMemo(n-2, retList)
	}
	return x
}

func getPiboWithGoRootin(n uint) uint {
	var x uint
	if 0 <= n && n <= 1 {
		x = n
	} else {
		x = getPibo(n-1) + getPibo(n-2)
	}
	return x
}

func getPibo(n uint) uint {
	var x uint
	if 0 <= n && n <= 1 {
		x = n
	} else {
		x = getPibo(n-1) + getPibo(n-2)
	}
	return x
}
