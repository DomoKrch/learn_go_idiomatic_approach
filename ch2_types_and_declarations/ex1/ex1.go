package main

import "fmt"

func main() {
	var i int = 20
	// Illegal without type conversion
	var f float64 = float64(i)

	fmt.Println(i)
	fmt.Println(f)
}
