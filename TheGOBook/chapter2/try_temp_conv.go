package main

import "github.com/crazcalm/learn-go/TheGOBook/chapter2/tempconv"
import "fmt"

func main(){
    c := tempconv.Celsius(10)
    f := tempconv.Fahrenheit(20)
    k := tempconv.Kelvin(0)
    fmt.Println(c.String())
    fmt.Println(f.String())
    fmt.Println(k.String())
    fmt.Println(tempconv.CToF(c))
    fmt.Println(tempconv.CToK(c))
    fmt.Println(tempconv.FToC(f))
    fmt.Println(tempconv.FToK(f))
    fmt.Println(tempconv.KToC(k))
    fmt.Println(tempconv.KToF(k))
}
