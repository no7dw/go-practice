package main

import "os"

func main() {
  // ' !=  ""
  // f := os.Open("README.md") //build error: return two parameter
  f, err := os.Open("README.md") //build error: err not used
  
  defer f.Close() //will however run this line of code
  //read from f
}