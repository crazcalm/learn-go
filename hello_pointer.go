package main
import "fmt"

//Note that i1 and intP both point to the same value

func main(){
    var i1 = 5
    fmt.Printf("An integer %d, its location in memory: %p\n", i1, &i1)

    var intP *int // intP is a variable that pointer
    intP = &i1
    fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
