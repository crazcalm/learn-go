package main

import(
    "fmt"
    "strings"
)

func main(){
    var orig string = "Hey, how are you George?"
    var lower string
    var upper string
    fmt.Printf("The original string is: %s\n", orig)

    lower = strings.ToLower(orig)
    upper = strings.ToUpper(orig)

    fmt.Printf("The lowercase string is: %s\n", lower)
    fmt.Printf("The uppercase string is: %s\n", upper)
}
