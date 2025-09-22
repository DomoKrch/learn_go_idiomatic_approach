package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[1:2]

	// y subslice modified the x parent slice since they share the same memory in some points
	fmt.Println("Using classical subslices:")
	y = append(y, "e")
	fmt.Println(x, cap(x))
	fmt.Println(y, cap(y))

	// Full slice expression limits on how much the parent slice x shares memory with the
	// sublsice z. The resulting capacity of z = third offset showing the position where
	// subslice stops sharing memory with original slice x - the first offset which is
	// the starting index of subslice z in original slice x
	z := x[1:2:2]
	fmt.Println("\nUsing full slice expression for subslices:")
	z = append(z, "f")
	fmt.Println(x, cap(x))
	fmt.Println(z, cap(z))
}
