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
//change ref object ,so need to use pointer instead of struct object
func (e *Employee) rename(newName string) {
	e.name = newName
}
func main() {
	p := People{"440184"}
	e := Employee{People{"440183"},"dengwei", 32}
	e1 := Employee{p,"dengwei2", 32}
	fmt.Println(p,e,e1)
	fmt.Println(p.sayHi())
	fmt.Println(e.sayHi())
	fmt.Println(e.selfIntro())

	fmt.Println(e1.name)
	e1.rename("Wade")
	fmt.Println(e1.name)
}
