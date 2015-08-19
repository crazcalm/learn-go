package main

import "fmt"

type Shaper interface{
    Area() float32
}

type Square struct{
    side float32
}

func (sq *Square) Area() float32{
    return sq.side * sq.side
}

type Rectangle struct{
    length, width float32
}

func (r Rectangle) Area() float32{
    return r.length * r.width
}

func main(){
    r := Rectangle{5, 3}
    q := &Square{5}

    // An array of shapers
    // shapes := []Shaper{Shaper(r), Shaper(q)}
    // or shorter
    shapes := []Shaper{r, q}
    
    fmt.Println("Looping through shapes for area ...")

    for n, _ := range shapes{
        fmt.Println("Shape details: ", shapes[n])
        fmt.Println("Area of this shape is: ", shapes[n].Area())
    }
}
