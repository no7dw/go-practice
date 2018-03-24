package main

import "fmt"

type Employee struct {
	name string
	age  int
}

func main() {
	e := Employee{"dengwei", 32}
	fmt.Println(e)
}
