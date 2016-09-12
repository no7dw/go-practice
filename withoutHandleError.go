package main

import "os"

func main() {
  // ' !=  ""
  f, err := os.Open("README.md") //build error: err not used
  
  defer f.Close() //will however run this line of code
  //read from f
}