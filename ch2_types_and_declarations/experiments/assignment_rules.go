package main

import "fmt"


func main() {
	x := 30
	var z rune = 'f'

	// This breaks since lefthand variables include only x, i.e. no new variables
	// x := 40

	// But apparently, you can modify old existing var with :=, if lefthand variables include
	// something new (in this case, its y variable). this is pretty stupid, if you ask me
	x, y := 40, "hi"

	/// This is ok
	z = 'a'


	fmt.Println(x, y, string(z))
}
