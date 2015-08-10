//Mutually recursive: Functions can also be used in Go: These functions
//that call one another. Because of the Go compilation process, these
//functions may be declared in any other. Here is a simple example:
//even calls odd, and odd calls even

package main

import "fmt"

func main(){
    fmt.Printf("%d is even: is %t\n", 16, even(16))
    fmt.Printf("%d is odd: is %t\n", 17, odd(17))
    fmt.Printf("%d is odd: is %t\n", 18, odd(18))
}

func even(nr int) bool{
    if nr == 0 {
        return true
    }
    return odd(RevSign(nr) - 1)
}

func odd(nr int) bool{
    if nr == 0{
        return false
    }
    return even(RevSign(nr) - 1)
}

func RevSign(nr int)int{
    if nr < 0{
        return -nr
    }
    return nr
}
