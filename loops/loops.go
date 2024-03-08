package main

import "fmt"

func main() {
	fmt.Println("for loop")
	//for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("for loop modified as while loop")
	//for loop modified as while loop
	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}
}
