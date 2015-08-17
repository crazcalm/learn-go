package main

import "fmt"

type number struct{
    f float32
}

type nr number // alais type

func main(){
    a := number{5.0}
    b := nr{5.0}

    // This is the conversion...
    var c = number(b)

    fmt.Println(a,b,c)
}
