package main

import "fmt"

func sum(a, b int) int{
  return a + b
}

func main(){

  var answer = sum(1, 2)
  fmt.Println("The answer is ", answer)
}
