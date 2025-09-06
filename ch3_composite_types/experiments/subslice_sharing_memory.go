package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := x[:3]

	fmt.Println("\n1) Before modification:")
	// Both original slice (x) and its usbslice (y) still share the same capacity
	fmt.Println("x: ", x)
	fmt.Println("Capacity and length of x: ", cap(x), len(x))
	fmt.Println("y: ", y)
	fmt.Println("Capacity and length of y: ", cap(y), len(y))

	y = append(y, 5)

	fmt.Println("\n2) After appending 5 to the y:")
	// If subslice's element is modified, then original slice's element is modified as well
	fmt.Println("x: ", x)
	fmt.Println("Capacity and length of x: ", cap(x), len(x))
	fmt.Println("y: ", y)
	fmt.Println("Capacity and length of y: ", cap(y), len(y))

	x[3] = 0
	y = append(y, 6)

	fmt.Println("\n3) After appending 6 to the y and modifying x's element:")
	// y subslice doubled its capacity and created its own backing array (away from original slice - x)
	// original slice (x) seems to have saved modified data from subslice (before capacity overflow)
	// tested it. this seems to happen when the capacity overflow occurs with appending only one
	// overflowing element to the subslice (boogles the mind, i know)
	fmt.Println("x: ", x)
	fmt.Println("Capacity and length of x: ", cap(x), len(x))
	fmt.Println("y: ", y)
	fmt.Println("Capacity and length of y: ", cap(y), len(y))

	y[0] = 7
	y[1] = 8
	x[3] = 9

	fmt.Println("\n4) After trying to modify x and y:")
	// x and y have their own backing arrays
	fmt.Println("x: ", x)
	fmt.Println("Capacity and length of x: ", cap(x), len(x))
	fmt.Println("y: ", y)
	fmt.Println("Capacity and length of y: ", cap(y), len(y))

	z := y[:2]
	z = append(z, 11, 12, 13, 14, 15, 16, 17)
	z[0] = 18

	fmt.Println("\n5) After creaing a sublice (z) from y, appending elements over its capacity and trying to modify z's elements:")
	// subslice (z) has created its own backing array due to capacity oveflow, BUT
	// its original slice (y) didn't save any modified data. tested this too, this
	// happens only when I append A SEQUENCE OF ELEMENTS WITH THE OVERFLOWING ELEMENT TOO
	// (Go runtime ALLOCATES A NEW BACKING ARRAY FIRSTLY, moves old elements of z to it
	// and then adds new elements -> that's why y element's aren't touched)
	fmt.Println("x: ", x)
	fmt.Println("Capacity and length of x: ", cap(x), len(x))
	fmt.Println("y: ", y)
	fmt.Println("Capacity and length of y: ", cap(y), len(y))
	fmt.Println("z: ", z)
	fmt.Println("Capacity and length of z: ", cap(z), len(z))


}
