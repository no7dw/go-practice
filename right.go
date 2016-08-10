package main

import "fmt"

//var a := 0 //wrong
var a int = 0
//c := 1 //wrong
//d = 1 //wrong

func cal(b int)(val1 int, val2 int){
  fmt.Println(b)
  val1 = 1
  val2 = 2
  return 
}
func test(){
  c := 1
  //d = 1 //wrong
  var f , d int
  f,d = cal(1)
  fmt.Println(c, f, d)
}
func main() {
  fmt.Println("Hello, 世界")
}
