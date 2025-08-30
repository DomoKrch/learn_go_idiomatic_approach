package main

import "fmt"

func main() {
	const value = 100

	// Legal
	var i int = value
	var f float64 = value

	fmt.Println(i)
	fmt.Println(f)
}
