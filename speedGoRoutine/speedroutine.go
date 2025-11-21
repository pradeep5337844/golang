package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch:= make(chan string)
	cities:= []string{"bangalore", "mangalore", "kolar","yelahanka","devanahalli", "avathi"}
	starttime:= time.Now()
	for _,city := range cities {
		wg.Add(1)
		go cityname(city, ch, &wg)
	}
	go func(){
		wg.Wait()
		close(ch)
	}()

	for c := range ch{
		fmt.Println(c)
	}
	fmt.Println("execution took ", time.Since(starttime))
}

func cityname(city string, ch chan string, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(1*time.Second) 
	fmt.Println("the city is ", city)
	ch <- city
}
