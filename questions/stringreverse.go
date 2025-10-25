package main

import "fmt"

func main() {
	str := "golang"
	reversed := ""
	for _, ch := range str {
		reversed = string(ch) + reversed
	}
	fmt.Println("Reversed:", reversed)
}
