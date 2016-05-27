package main

import "C"

//export Hi
func Hi() string{
    hi := "hi"
    return hi
}

func main(){
    println(Hi())
}
