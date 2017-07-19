package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func square(x *float64) {
	*x = *x * *x
}

func main() {
	z := 0

	// How do you get the memory address of a variable?
	fmt.Println(&z)

	// How do you assign a value to a pointer?
	one(&z)
	fmt.Println(z)

	// How do you create a new pointer?
	y := new(int)
	fmt.Println(*y)

	// What is the value of x after running this program
	x := 1.5
	square(&x)
	fmt.Println(x)
}
