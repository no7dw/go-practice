package main

import (
	"fmt" // located in $GOROOT/src/
	"github.com/no7dw/go-practice/package/greet" // located in $GOPATH/src/
	"package/greet2" // located in $GOPATH/src/
)

func main() {
	fmt.Println(greet.Morning)
	fmt.Println(greet2.Morning)
}
