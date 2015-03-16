 package main

 import (
         "encoding/json"
         "fmt"
         "net/http"
 )

 type JsonResponse map[string]interface{}

 func (jr JsonResponse) String() (str string) { // add this method to print out String
         byte, err := json.Marshal(jr)
         if err != nil {
                 str = "" // return empty
                 return
         }
         str = string(byte) //ok,return cast byte to string
         return
 }

 func JsonReplyAdam(w http.ResponseWriter, r *http.Request) {
         w.Header().Set("Content-Type", "application/json")
         fmt.Fprint(w, JsonResponse{"name": "Adam", "age": 36, "job": "CEO", "success": true})
 }


 func JsonReplyEve(w http.ResponseWriter, r *http.Request) {
         w.Header().Set("Content-Type", "application/json")
         fmt.Fprint(w, JsonResponse{"name": "Eve", "age": 30, "job": "CFO", "success": true})
 }

 func main() {
         // http.Handler

         mux := http.NewServeMux()
         mux.HandleFunc("/adam", JsonReplyAdam)
         mux.HandleFunc("/eve", JsonReplyEve)

         http.ListenAndServe(":8000", mux)
 }