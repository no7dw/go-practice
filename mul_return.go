package main

import (
	"fmt"
	"errors"
)


//return float and standard error
func holesPerGopher(gophers, holes int)(float32 , error) {
	if gophers==0{
		return 0, errors.New("no gophers! not allow divide by zero");
	}
	return float32(holes)/float32(gophers), nil
	
}

func main() {
	gopher := 5
	gopherHoles := 2

	hpg, _ := holesPerGopher(gopher, gopherHoles);
	fmt.Println("Holes per gopher: ", hpg);
	//fmt.Println("Holes per gopher: ", hpg, _);  cannot use _ as value
}