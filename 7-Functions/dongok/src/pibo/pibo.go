package pibo

import (
	"fmt"
	"math/big"
	"reflect"
)

func bigIntSort(unsorted []big.Int) []big.Int {
	var tmp big.Int

	fmt.Println("notsorted", unsorted, reflect.TypeOf(unsorted))
	unsorted2 := unsorted

	for i := 0; i < (len(unsorted2) - 1); i++ {
		for j := i + 1; j < len(unsorted2); j++ {
			if unsorted2[i].Cmp(&unsorted2[j]) == 1 {
				tmp = unsorted2[j]
				unsorted2[j] = unsorted2[i]
				unsorted2[i] = tmp

			}
		}
	}

	fmt.Println("big.IntSort", unsorted)
	fmt.Println("big.IntSort2", unsorted2)
	return unsorted
}

/* GetRecursivePiboList
재귀형 피보나치 수열 구하는법
*/
func GetRecursivePiboList(n int) []big.Int {
	var pibolist []big.Int

	for i := 0; i <= n; i++ {
		a := big.NewInt(0)
		a = a.SetInt64(int64(i))
		getPiboV := getPibo(a)
		//fmt.Println("GetPibo", getPiboV)
		pibolist = append(pibolist, *getPiboV)
	}
	return pibolist
}

func GetMemoriPiboList(n int) []big.Int {
	fmt.Println("GetMemoriPiboList")
	var pibolist []big.Int
	retList := make(map[int64]big.Int)
	var v = big.NewInt(0)
	var i int
	for i = 0; i <= n; i++ {
		a := big.NewInt(0)
		a = a.SetInt64(int64(i))
		v = getPiboWithMemo(a, retList)
		pibolist = append(pibolist, *v)
		retList[int64(i)] = *v
		//fmt.Println(retList)
	}
	/*
		var keys []big.Int
		for k := range retList {
			keys = append(keys, k)
		}

		fmt.Println("a", keys)
		big.IntSort(keys)
		fmt.Println("b", keys)

		for k := range keys {
			v := retList[big.Int(k)]
			fmt.Println("key", k, "value", v)
		}
	*/
	return pibolist
}

func getPiboWithMemo(n *big.Int, retList map[int64]big.Int) *big.Int {
	value, ok := retList[n.Int64()]
	//fmt.Println("n", n, "retList", retList, "ok", ok)
	if ok {
		//fmt.Println("exist value", n)
		return &value
	}
	x := big.NewInt(0)
	if n.Cmp(big.NewInt(1)) < 1 {
		x = n
	} else {
		//fmt.Println("x", x)
		//fmt.Println("n", n)
		a := big.NewInt(0).Set(n)
		b := big.NewInt(0).Set(n)
		a.Sub(a, big.NewInt(1))
		//fmt.Println("a", a)
		b.Sub(b, big.NewInt(2))
		//fmt.Println("b", b)
		x.Add(getPiboWithMemo(a, retList), getPiboWithMemo(b, retList))
		//fmt.Println("x2", x)
	}
	return x
}

func getPibo(n *big.Int) (x *big.Int) {
	x = big.NewInt(0)
	if n.Cmp(big.NewInt(1)) < 1 {
		x = n
	} else {
		//fmt.Println("x", x)
		//fmt.Println("n", n)
		a := big.NewInt(0).Set(n)
		b := big.NewInt(0).Set(n)
		a.Sub(a, big.NewInt(1))
		//fmt.Println("a", a)
		b.Sub(b, big.NewInt(2))
		//fmt.Println("b", b)
		x.Add(getPibo(a), getPibo(b))
		//fmt.Println("x2", x)
	}
	return x
}
