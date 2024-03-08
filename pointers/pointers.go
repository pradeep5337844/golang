package main

import "fmt"

func main() {
	x := 5
	fmt.Println("value of variable x is:", x)
	fmt.Println("address of variable x is:", &x)
	changeValue(x)
	fmt.Println("the value of x with x as a variable is", x)
	//pass address of x
	changeValuePointer(&x)
	fmt.Println("the value of x with x as a pointer is", x)
}

func changeValue(x int) {
	x = 10
}

func changeValuePointer(x *int) {
	//modify the pointer variable
	*x = 10
}
