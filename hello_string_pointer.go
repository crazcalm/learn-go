package main
import(
    "fmt"
)

//Note that changing the value of the pointer also changes the
//value of the original variable

func main(){
    s := "good bye"
    var p *string = &s
    *p = "ciao"
    fmt.Printf("Here is the pointer p: %p\n", p)
    fmt.Printf("Here is the string *p: %s\n", *p)
    fmt.Printf("Here is the string s: %s\n", s)
}

