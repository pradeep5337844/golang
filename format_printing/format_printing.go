package main

import "fmt"

func main() {
	var name string = "pradeep"
	var a float32 = 20.13
	c := false
	b := 10
	fmt.Println("length of name is:", len(name))
	fmt.Println(name + " is an engineer")
	fmt.Printf("a is %f \n", a)
	fmt.Printf("a formatted to 2 decimal point is %.2f \n", a)
	fmt.Printf("type of name is %T \n", name)
	fmt.Printf("value of c is %t \n", c)
	fmt.Printf("type of name is %T \n", c)
	fmt.Printf("b is %d \n", b)
	fmt.Printf("binary of 25 is %b \n", 25)
}
