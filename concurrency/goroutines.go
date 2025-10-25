package main

import (
	"fmt"
	"time"
)

func main() {
	go iterate()
	time.Sleep(time.Second * 3)
}

func iterate() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}
