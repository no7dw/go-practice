package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//Name Age should be uppercase
//https://golang.org/ref/spec#Exported_identifiers
type Employee struct {
	Name string
	Age  int
}

func getDB(uri string) *mgo.Database {
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	c := session.DB("test")
	return c
}
func main() {
	db := getDB("localhost:27017")
	c := db.C("employee")
	result := Employee{}
	err := c.Find(bson.M{"name": "wade"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
