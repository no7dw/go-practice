package main

import (
	// "io"
	"net/http"
	"os"
	"fmt"
	"log"
	// "encoding/json"
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
	ip := string(data[:count-1])
	fmt.Printf("read : %s %d\n", ip, count)
	defer file.Close()
	return "{ \"ip\":\"" + ip + "\"}" 

}
func hello(w http.ResponseWriter, r *http.Request) {
	ip := read()
	w.Header().Set("Content-Type", "application/json")	
	fmt.Fprint(w, ip)
}

func main() {
	
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
