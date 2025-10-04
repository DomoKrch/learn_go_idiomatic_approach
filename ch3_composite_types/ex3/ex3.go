package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName string
		id int
	}

	xEmployee := Employee{
		"John",
		"Wick",
		1,
	}

	yEmployee := Employee{
		firstName: "Bob",
		lastName: "Dylan",
		id: 2,
	}

	var zEmployee struct {
		firstName string
		lastName string
		id int
	}

	zEmployee.firstName = "Frank"
	zEmployee.lastName = "Sinatra"
	zEmployee.id = 3

	fmt.Println(xEmployee, yEmployee, zEmployee)
}
