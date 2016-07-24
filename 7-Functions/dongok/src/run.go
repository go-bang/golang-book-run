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

	fmt.Println("pibolist", pibo.GetMemoriPiboList(uint(piboMax)))
	//fmt.Println("pibolist", pibo.GetRecursivePiboList(piboMax))
}
