package main

import (
	"fmt"
)

/*
Head  is saa
dfasf
*/
type Head struct {
	eye string
}

func (x *Head) see() {
	fmt.Println(x.eye)
}

/*
Body fdsfad
*/
type Body struct {
	h Head
	Head
}

func main() {
	//var h Head
	h := Head{"aaa"}
	fmt.Println(h)
	h.see()

	var b Body
	fmt.Println("b", b)
	b.h.see()
	b.see()
}
