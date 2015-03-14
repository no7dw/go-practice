package main

import "os"

func main() {
	// ' !=  ""
	f, err := os.Open("README.md")
	if err!= nil{
		return
	}
	defer f.Close() //will however run this line of code
	//read from f
}