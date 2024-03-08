package main

import "fmt"

func main() {
	var a int
	var b, l bool
	var c float32
	var (
		m = 5
		n = 6
	)
	a = 10
	b = false
	l = true
	c = 10.20
	x, y := 15, 16

	fmt.Println("x+y:", x+y)
	fmt.Println("n-m::", n-m)
	fmt.Println("a*c:", float32(a)*c)
	fmt.Println("x/m:", x/m)
	fmt.Println("y%m:", y%m)
	fmt.Println("b is :", b)
	if !b {
		fmt.Println("not operation")
	}
	k := x + y
	if b || k == 31 {
		fmt.Println("or operarion")
		if l && k == 31 {
			fmt.Println("and operation")
		}
	}

}
