package main

import "fmt"

// If you want to change the value of the an array,
// you must pass a pointer to the array.

func f(a [3]int){ fmt.Println(a) }

func fp(a *[3]int){ fmt.Println(a) }

func main(){
    var ar [3]int

    f(ar)// passes a copy of ar
    fp(&ar)// passes a pointer to ar
}
