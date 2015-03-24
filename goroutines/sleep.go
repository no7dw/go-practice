package main
import (
    "fmt";
    "time"
)
func IsReady(what string, minutes int) {
	time.Sleep(time.Duration(minutes) * time.Second )
    // time.Sleep(minutes * 60*1e9)// Unit is nanosecs.
    fmt.Println(what, "is ready")
}
func main(){
    go IsReady("tea", 6)
    go IsReady("coffee", 2)
    fmt.Println("I'm waiting…")
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}
