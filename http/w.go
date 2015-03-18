package main

import (
    "os"
    "fmt"
)

func main() {
    f, err := os.Open("testfile")
    if err != nil {
        fmt.Printf("Error: Can't create config file, %s\n", err)
        return
    }
    defer f.Close()
    
    str := "some string"
    
    n, err := f.WriteString(str)
    if n != len(str) {
        if err == nil {
            panic("err ought to be something")
        }
        fmt.Printf("Error: Write has failed: %s\n", err) //this is the output, not the panic
        return
    }
}