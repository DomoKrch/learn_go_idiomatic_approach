package main

import "fmt"

func main() {
	// Non-nil map with length of 0
	xMap := map[string]int{}

	// Nil map
	var yMap map[string]int

	// Non-nil map with length of 0 and capacity of 10
	zMap := make(map[string]int, 10)

	// This would return 0, even if the key didn't exist before
	fmt.Println("\nxMap:")
	fmt.Println(xMap["test"])
	xMap["test"]++
	// This has been really incremented and its going to return 1
	fmt.Println(xMap["test"])

	// This will return 0 when trying to read non-existen key, but you can't modify it since this map equals to nil
	fmt.Println("\nyMap:")
	fmt.Println(yMap["another_test"])
	// yMap["another_test"]++
	// fmt.Println(yMap["another_test"])

	// Same situation as with xMap
	fmt.Println("\nzMap:")
	fmt.Println(zMap["one_more_test"])
	zMap["one_more_test"]++
	fmt.Println(zMap["one_more_test"])

}
