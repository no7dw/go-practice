package main
import "fmt"
import "time"
func f(from string) {
    for i := 0; i < 3; i++ {
	time.Sleep(1 * time.Second )
        fmt.Println(from, ":", i)
    }
}
func main() {

    f("direct")

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}
