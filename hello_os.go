package main

import "fmt"
import "os"

func main(){

    test, _ := os.Getwd()
    
    fmt.Printf(test)
}
