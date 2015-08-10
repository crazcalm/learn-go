package main

import "fmt"

func main(){
    var min, max int
    min, max = MinMax(78, 68)
    fmt.Printf("Minimum is: %d, Maximum is: %d\n", min, max)
}

func MinMax(a, b int)(min, max int){
    if a<b{
        min = a
        max = b
    }else{
        min = b
        max = a
    }

    return
}
