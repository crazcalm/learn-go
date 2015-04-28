package main

import(

"fmt"
"sort"
)


func main(){

    test := [] int {4,1,3,2,5,2,9}
    sort.Ints(test)
    fmt.Println(test)
    
}
