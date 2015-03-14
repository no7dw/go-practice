package main

import (
	"fmt"
	"time"
)



func main(){
	// list :=[5]int {1,2,3,4,5}
	c := make(chan int) //allocate a channel

	go func() {
		// list.Sort()
		time.Sleep(time.Duration(2) * time.Second)	
		fmt.Println("Sort complete, send signal")
		c <- 1 //send a signal, whatever value
	}()

	//do somethins for a while
	time.Sleep(time.Duration(1) * time.Second)
	//try time.Sleep(time.Duration(3) * time.Second)
	fmt.Println("wait...")
	<-c //wait for sort to finish , discard send value	
	fmt.Println("done")
}