package main

import "fmt"

func main() {
	var message string = "Hi 👩 and 👨"
	// Instead of manually traversing string's bytes and risking getting in the middle of byte collection for some
	// emoji, i am converting message to its rune (int32) representation in order to safely capture the full byte
	// collection for emoji character in string
	runeMessage := []rune(message)

	fmt.Println(string(runeMessage[3]))
}
