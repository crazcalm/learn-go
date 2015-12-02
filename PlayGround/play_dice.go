package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// creating and seeding rand generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// map counter for a die
	die := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0}

	// simulate rolling the die
	for i := 0; i < 36; i++ {
		die[r.Intn(6)]++
	}

	// print results to screen
	for i := 0; i < 6; i++ {
		fmt.Printf("%d:%d\n", i, die[i])
	}

}
