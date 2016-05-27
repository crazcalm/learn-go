package main

import (
    "C"
)

//export Greetings
func Greetings(name string) string{
    return "hello " + name
}

func main(){
    name := "marcus"
    println(Greetings(name))
}
