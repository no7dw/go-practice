// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"log"
	"encoding/json"
	// "math/rand"
	"net/url"
	"os"
	"os/signal"
	// "time"
	// "fmt"
	"github.com/gorilla/websocket"
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
var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			mt, message, err := c.ReadMessage()
			var emp OrderBook
			err = json.Unmarshal(message, &emp)
			log.Println("read:",mt, err, emp)
			if err != nil {
				return
			}
			// log.Printf("recv: %s", string(message[:]))
		}
	}()

	// ticker := time.NewTicker(time.Second)
	
	// defer ticker.Stop()
	for {
		// mt, message, err := c.ReadMessage()
		// if err != nil {
		// 	log.Println("read2", err)
		// 	break
		// }
		// log.Println("message",mt, string(message[:]))
		select {
		// case <-done:
		// 	return
		// case t := <-ticker.C:
		// // case	t := rand.ExpFloat64():
		// 	// s := fmt.Sprintf("%f", t)
		// 	// err := c.WriteMessage(websocket.TextMessage, []byte(s))
		// 	err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))

			
		// 	if err != nil {
		// 		log.Println("write:", err)
		// 		return
		// 	}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
		// 	select {
		// 	case <-done:
		// 	case <-time.After(time.Second):
		// 	}
		// 	return
		}
	}
}
