package main

import "fmt"

func main(){
    var arrAge = [5]int{18, 20, 15, 22, 16}
    var arrLazy = [...]int{5, 6, 7, 8, 22}// Note: size not given

    var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
    
    for i:=0; i<len(arrKeyValue); i++{
        fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
    }
    fmt.Println()

    for i:=0; i<len(arrAge); i++{
        fmt.Printf("arrAge[%d] = %d\n", i, arrAge[i])
    }
    fmt.Println()

    for i:=0; i<len(arrLazy); i++{
        fmt.Printf("arrLazy[%d] = %d\n", i, arrLazy[i])
    }
    
}
