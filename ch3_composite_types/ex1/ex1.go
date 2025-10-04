package main

import "fmt"

func main() {
	var xSlice = []string{"Hello", "Hola", "नमस्ते", "こんにちは", "Привіт"}

	xSubSlice := xSlice[:2]
	ySubSlice := xSlice[1:4]
	zSubSlice := xSlice[3:]

	fmt.Println(xSlice, xSubSlice, ySubSlice, zSubSlice)
}
