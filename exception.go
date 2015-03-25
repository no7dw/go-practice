package main
import "fmt"

func main() {
 defer func() {
   if r := recover(); r != nil {
     fmt.Println("recover", r)
} }()
 someFunc() 
}
func someFunc() {
 panic("[fail]")
}
