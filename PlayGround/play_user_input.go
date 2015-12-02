package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string

	input := bufio.NewScanner(os.Stdin)

	fmt.Println("What is your name?")

	for input.Scan() {

		name = input.Text()
		if name != "" {
			break
		}

		fmt.Println("That name is no good. Try again.")
	}

	// Note: ignoring potential errors from input.Err()
	// Note: Don't know how to deal with those errors...
	fmt.Printf("Hello %s\n", name)
}
