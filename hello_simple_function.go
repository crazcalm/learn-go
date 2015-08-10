package main

import "fmt"

func main(){
    fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
}

func MultiPly3Nums(a, b, c int) int {
    return a*b*c
}
