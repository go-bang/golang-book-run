package main

import (
	"fmt"
	htkim_math "go-bang/golang-book-run/11-Packages/htkim/math"
)

func main() {
	l := []float64{1, 2, 3}
	fmt.Println("min: ", htkim_math.Min(l))
	fmt.Println("avg: ", htkim_math.Average(l))
	fmt.Println("max: ", htkim_math.Max(l))
}



