package main

import "fmt"

func main() {
	var x int64 = 20
	var y byte = 23
	 //z := "a"

	// var baz float64 = 1000.23
	// This fails, makes sense, int < float64, just wanted to check
	// var qux int = baz


	// No automatic type promotion, everything has to be explicitly converted to a particular type
	// var foo string = 63

	// This works though
	var bar int = 'a'

	fmt.Println(byte(x) + y)

	// Sometimes it doesn't work. Can't convert string to int64 via type conversion technique.
	// However, it would work with integer
	// fmt.Println(string(z) + x)

	// fmt.Println(qux)

	// fmt.Println(foo)
	fmt.Println(bar)
}
