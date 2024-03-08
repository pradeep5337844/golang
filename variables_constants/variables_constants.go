package main

import (
	"fmt"
)

func main() {
	//constant is a variable whose value doesnt change
	const pi float64 = 3.1415

	//variable is the one whose values can be changed
	var (
		a         = 10
		b string  = "hello"
		c float32 = 0.112
	)
	fmt.Println("constant pi value is: ", pi)
	fmt.Println("the values of a b and c before changing values", a, b, c)
	a = 20
	b = "world"
	c = 1.23
	fmt.Println("the values of a b and c after changing values", a, b, c)
	fmt.Printf("the type of a is %T\n", a)
	fmt.Printf("the type of b is %T\n", b)
	fmt.Printf("the type of c is %T\n", c)
	fmt.Printf("the type of pi is %T\n", pi)
}
