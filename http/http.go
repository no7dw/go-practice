package main

import (
    // "io"
    "net/http"
    "os"
    "fmt"
    "log"
    "encoding/json"
    // "strings"
    "io/ioutil"  
)
type Server struct {
    IP string
}
func read()(string) {
    file, err := os.Open("ip.log") // For read access.
    if err != nil {
        log.Fatal(err)
    }   
    data := make([] byte , 16)  
    count, err := file.Read(data)
    if err != nil {
        log.Fatal(err)
    }
    ip := string(data[:count])
    fmt.Printf("read : %s %d\n", ip, count)
    defer file.Close()
    return "{ \"IP\":\"" + ip + "\"}" 

}
func getip(w http.ResponseWriter, r *http.Request) {        
    //ip := read()
    w.Header().Set("Content-Type", "application/json")  
    fmt.Fprint(w, "123.222.222.111")
}
func iphandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET"{
        getip(w,r)
    }else if r.Method == "POST"{
        result, _:= ioutil.ReadAll(r.Body)  
        r.Body.Close() 
        // log.Println(string(result))         
        // fmt.Println(result)        
        var s Server;  
        json.Unmarshal([]byte(result), &s) 

        fmt.Println("save ip:",s.IP)
        err := os.Remove("ip.log")
        f, err := os.OpenFile("ip.log",  os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Printf("Error: Can't create config file, %s\n", err)
            return
        }
        defer f.Close()
         
        str := s.IP        
        
        n, err := f.WriteString(str)
        if n != len(str) {
            if err == nil {
                panic("err ought to be something")
            }
            fmt.Printf("Error: Write has failed: %s\n", err) //this is the output, not the panic
            return
        }
       
     }
}

func main() {
    
    http.HandleFunc("/", getip)
    http.HandleFunc("/ip", iphandler)
    http.ListenAndServe(":8000", nil)
}
