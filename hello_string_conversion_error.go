package main

import(
    "fmt"
    "strconv"
)

func main(){
    var orig string = "ABC"
    var an int
    var err error

    an, err = strconv.Atoi(orig)

    if err != nil {
        fmt.Printf("orig %s is not a integer-existing with error\n", orig)
        return
    }

    fmt.Printf("The iteger is %d\n", an)

    //rest of the code
}
