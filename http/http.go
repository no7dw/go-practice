package main

import (
	// "io"
	"net/http"
	"os"
	"fmt"
	"log"
	"encoding/json"
	
    "io/ioutil"  
)
type Server struct {
	ServersIP string
}
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
func getip(w http.ResponseWriter, r *http.Request) {
	ip := read()
	w.Header().Set("Content-Type", "application/json")	
	fmt.Fprint(w, ip)
}
func iphandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		getip(w,r)
	}else if r.Method == "POST"{
		result, _:= ioutil.ReadAll(r.Body)  
        r.Body.Close()  
        fmt.Println(result)        
        var s Server;  
        json.Unmarshal([]byte(result), &s)  
        file, err := os.Open("ip.log")
        if err != nil {
        	fmt.Println("open")
			log.Fatal(err)
		}
		count, errw := file.WriteString(s.ServersIP)
		if errw != nil {
			fmt.Println("WriteString")
			log.Fatal(errw)
		}
		fmt.Println("write: %d",count);
		defer file.Close()        
        fmt.Println(s.ServersIP);  
     }
}

func main() {
	
	http.HandleFunc("/", getip)
	http.HandleFunc("/ip", iphandler)
	http.ListenAndServe(":8000", nil)
}
