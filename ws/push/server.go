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
	"time"
	"strconv"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	defer c.Close()
	counter := 1
	for {
		// mt, message, err := c.ReadMessage()
		// if err != nil {
		// 	log.Println("read:", err)
		// 	break
		// }
		// log.Printf("recv: %s", message)
		message := "abc"
		mt := 1 //TextMessage = 1
		select {
			case <-ticker.C:
				message += strconv.Itoa(counter)
				counter++
				err = c.WriteMessage(mt, []byte(message))
				if err != nil {
					log.Println("write:", err)
					break
				}
				log.Println("write:" , message)
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	// http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

