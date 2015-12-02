package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// seed/create random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 20; i++ {
		fmt.Println(r.Intn(6))
	}
}
