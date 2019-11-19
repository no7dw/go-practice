// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	// "encoding/binary"
	"encoding/json"
	// "bytes"
	// "strconv"
)
type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

 
type OrderBook struct {
	Asks	[][]string `json:"asks"`
	Bids [][]string `json:"bids"`
	Pair string `json:"pair"`
}
func getDB(uri string) *mgo.Database {
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	c := session.DB("nowdone")
	return c
}

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options


func find(dbc mgo.Database) (OrderBook , error){
	c := dbc.C("okex_order_books")
	// result := Employee{}
	// err := c.Find(bson.M{"name": "wade"}).One(&result)
	result := OrderBook{}
	err := c.Find(bson.M{}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func getEmpFromDB() OrderBook{
	
	emp, err := find(*db)
	if err != nil {
		log.Println(emp)
	}
	// log.Println(emp,err)
	return emp
}
func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	defer c.Close()
	// counter := 1
	for {
		// mt, message, err := c.ReadMessage()
		// if err != nil {
		// 	log.Println("read:", err)
		// 	break
		// }
		// log.Printf("recv: %s", message)
		
		// message := "abc"
		

		mt := 2 //TextMessage = 1 BinaryMessage = 2
		select {
			case <-ticker.C:
				// message := Employee{"dengwei", 32}
				// message += strconv.Itoa(counter)
				// counter++
				message := getEmpFromDB()
				// var bin_buf bytes.Buffer
				// binary.Write(&bin_buf, binary.LittleEndian, message)
				bin_buf, err := json.Marshal(message)
				err = c.WriteMessage(mt, bin_buf)
				if err != nil {
					log.Println("write:" , err)
					break
				}
				log.Println("write:" , message)
		}
	}
}

var db *mgo.Database

func main() {
	db = getDB("username:password@10.10.20.85:3717/nowdone")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	// http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

