package main

import(
    "fmt"
)

func main(){
    var num1 int = 99

    switch num1{
        case 98, 99:
            fmt.Println("It's equal to 98 or 99")

        case 100:
            fmt.Println("It's equal to 100")

        default:
            fmt.Println("It's not equal to 98, 99, or 100")
    }
}
