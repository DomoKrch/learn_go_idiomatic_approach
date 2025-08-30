package main

import "fmt"

func main() {
	const x int = 100
	// y := 10
	// z := 20

	// Breaks since constants are immutable and are only evaluated at compile time
	// x = z + y
	// x = x + 1

	fmt.Println(x)
}
