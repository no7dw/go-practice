package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)
//Name Age should be uppercase
type Employee struct {
	Name string 
	Age  int 
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("test").C("employee")
	result := Employee{}
	err = c.Find(bson.M{"name": "wade"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println( result)
}
