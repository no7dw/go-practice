package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Employee struct {
	name string
	age  int
}

func main() {
	fmt.Println("vim-go")
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("test").C("employee")
	result := Employee{}
	err = c.Find(bson.M{"name": "dengwei"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dengwei age: ", result)
}
