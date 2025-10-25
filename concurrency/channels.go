package main

import (
	"fmt"
	"sync"
)
func main(){
	var wg sync.WaitGroup
	ch:= make(chan int)
	go printnum(ch,&wg)
	wg.Add(1)
	for i:=0;i<100;i++{
		
		ch <- i
	}
	
	close(ch)
	wg.Wait()
}

func printnum(ch <-chan int,wg *sync.WaitGroup){
	defer wg.Done()
	for v := range ch {
		
        fmt.Println(v)
    }
}