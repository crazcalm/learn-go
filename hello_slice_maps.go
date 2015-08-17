package main

import (
    "fmt"
)

func main(){
    // Version A: Correct way. I am ignoring the incorrect way...
    items := make([]map[int]int, 5)

    for i:= range items{
        items[i] = make(map[int]int, 1)
        items[i][1] = 2
    }

    fmt.Printf("Version A: Value of items: %v\n", items)
}
