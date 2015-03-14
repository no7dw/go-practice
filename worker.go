package main

import (
	"fmt"
	"time"
)

func worker(workernum int, c <-chan int) {
	for _ = range c{
		fmt.Printf("%d got work\n", workernum)
		time.Sleep(time.Duration(1) * time.Second)
	}
	
}

func main(){
	// list :=[5]int {1,2,3,4,5}
	c := make(chan int) //allocate a channel
	//try for i := 0; i < 1; i++ {
	for i := 0; i < 4; i++ {
		go worker(i, c)
	}

	for i := 0; i < 10; i++ {
		c <- i //block until all is read
	}
	fmt.Println("read all channel done")
}