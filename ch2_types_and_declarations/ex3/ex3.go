package main

import "fmt"

func main() {
	var b byte = 255 // 2^n - 1
	var smallI int32 = 2147483647 // 2^(n-1) - 1
	var bigI uint64 = 18446744073709551615 // 2^n - 1

	fmt.Println("Before adding 1:")
	fmt.Println("b: ", b)
	fmt.Println("smallI: ", smallI)
	fmt.Println("bigI: ", bigI)

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	// Overflow
	fmt.Println("After adding 1:")
	fmt.Println("b: ", b)
	fmt.Println("smallI: ", smallI)
	fmt.Println("bigI: ", bigI)
}
