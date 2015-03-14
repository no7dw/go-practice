package main

import (
	"io"
	"net/http"
	"os"
	"fmt"
	"log"
)
func read()(string) {
	file, err := os.Open("ip.log") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	data := make([] byte , 20)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	ip := string(data)
	fmt.Printf("read : ", ip, count)
	return ip 

}
func hello(w http.ResponseWriter, r *http.Request) {
	ip := read()
	io.WriteString(w, ip)//why is keep download a file?
}

func main() {
	
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
