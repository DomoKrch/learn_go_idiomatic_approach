package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := x[1:3]
	z := x[:4]


	// Capacity for subslice: Capacity of original slice - the starting offset of subslice in originla slice
	fmt.Println(x, cap(x))
	fmt.Println(y, cap(y))
	fmt.Println(z, cap(z))
}
