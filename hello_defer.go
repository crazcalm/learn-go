package main

import "fmt"

func main(){
    Function()
}

func Function(){
    fmt.Printf("In Function1 at the top\n")
    defer Function2()
    fmt.Printf("In Function1 at the bottom!\n")
}

func Function2(){
    fmt.Printf("Function2: Defered until the end of the calling function!")
}
