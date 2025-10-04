package main

import "fmt"

func main() {
	type car struct {
		colour string
		age int
	}


	nissan := car{
		colour: "blue",
		age: 4,
	}

	bmw := car{
		colour: "black",
		age: 2,
	}

	honda := car{
		colour: "blue",
		age: 4,
	}

	// anonymous struct
	dog := struct {
		colour string
		age int
	}{
		colour: "black",
		age: 2,
	}

	// false, since fields in structs differ
	fmt.Println(nissan == bmw)
	// true
	fmt.Println(nissan == honda)

	// true, anonynous struct dont need type conversion to be compared with another struct type (if
	// field ordering, field types and field names are the same)
	fmt.Println(bmw == dog)
}
