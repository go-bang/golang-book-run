package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	var x [5]int64

	x[0] = 12
	x[1] = 45
	x[2] = 34

	var total int64
	total = 0

	var avg float64

	for i := 0; i < 5; i++ {
		total = total + x[i]
	}

	fmt.Println(total)

	avg = float64(total / int64(len(x)))
	fmt.Println(avg)

	for _, v := range x {
		fmt.Println(v)
	}

	var slice1 []string

	slice1 = make([]string, 10)

	slice1[0] = "aaaa"

	slice2 := make([]string, 5)
	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(slice2))

	fmt.Println(len(slice2))
	for _, x := range slice2 {
		x = "1"
		fmt.Println(x)
	}

	for j := 0; j < len(slice2); j++ {
		slice2[j] = string(j)
	}
	slice2 = append(slice2, "0", "2")
	slice2 = append(slice2, "3", "4")
	fmt.Println(len(slice2))
	fmt.Println(slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice3, slice4)
	fmt.Println("print slice3, slice4", slice3, slice4)
	slice3[2] = 123
	fmt.Println("print slice3, slice4", slice3, slice4)
	//map

	var km = make(map[string]int)
	km["s1"] = 1
	fmt.Println(km, len(km))

	ki := make(map[int]int)
	ki[1] = 10
	ki[0] = 20
	fmt.Println(ki, ki[1], len(ki))

	delete(ki, 1)
	fmt.Println(ki, ki[1], len(ki))

	g := make([]int, 3, 5)
	fmt.Println("g", g, len(g))

	g1 := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(g1[2:5], g1[1:5])

	sample := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	sort.Ints(sample)

	fmt.Println(sample)

}
