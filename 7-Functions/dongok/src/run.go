package main

import (
	"fmt"
	pibo "pibo"
)

func main() {
	fmt.Println("go main")

	var piboMax int
	fmt.Print("input pibonachi length =")
	fmt.Scanf("%d", &piboMax)

	//fmt.Println("pibolist", pibo.GetMemoriPiboList(uint(piboMax)))
	/*
		for v, i := range pibo.GetRecursivePiboList(piboMax) {
			fmt.Println("index", v, "value", i.String())
		}
	*/
	for v, i := range pibo.GetMemoriPiboList(piboMax) {
		fmt.Println("index", v, "value", i.String())
	}
}
