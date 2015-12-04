package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	days := flag.Int("days", 10000, "Number of days to pass")
	prisoners := flag.Int("prisoners", 100, "Number of prisoners")

	flag.Parse()

	prisonersDilemma(days, prisoners)
}

// Need to review pointers!
func prisonersDilemma(days, prisoners *int) {
	// Create and seed random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Creating cells for prisoners
	cells := make(map[int]int)

	for i := 0; i < *prisoners; i++ {
		cells[i] = 0
	}

	// Simulate prisoners getting called to the room
	for i := 0; i < *days; i++ {
		cells[r.Intn(*prisoners)]++
	}

	// Print results to screen
	fmt.Printf("%d days have past\n\n", *days)
	for i := 0; i < *prisoners; i++ {
		fmt.Printf("Cell Number: %d: %d\n", i, cells[i])
	}
}
