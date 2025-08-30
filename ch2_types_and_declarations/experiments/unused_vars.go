package main

import "fmt"

func main() {
	// assignment not read, still ok
	x := 10

	// Breaks since y isn't read EVEN once -> compile-time error
	// var y rune = '\\'

	x = 20
	fmt.Println(x)

	// assignment not read, still ok
	x = 30
}
