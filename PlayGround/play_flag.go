package main

import (
	"flag"
	"fmt"
)

func main() {
	saying := flag.String("saying", "Stranger", "msg added to hello")

	flag.Parse()

	fmt.Println("Hello ", *saying)
}
