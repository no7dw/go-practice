package main

import (
	"fmt"
 	"strconv"
)
 
type People struct {
	idcard string
}

type Employee struct {
	people People
	name string
	age  int
}

func (p People) sayHi() string{
	return "hello"
}

func (e Employee) sayHi() string {
	return "hi"
}

func (e Employee) selfIntro() string {
	return strconv.Itoa(e.age)
}

func main() {
	p := People{"440184"}
	e := Employee{People{"440183"},"dengwei", 32}
	e1 := Employee{p,"dengwei2", 32}
	fmt.Println(p,e,e1)
	fmt.Println(p.sayHi())
	fmt.Println(e.sayHi())
	fmt.Println(e.selfIntro())
}
