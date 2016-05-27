package main

import "C"

//export HelloSlice
func HelloSlice() []int {
	slice1 := make([]int, 5)

	slice1[0] = 5
	slice1[1] = 1
	slice1[2] = 4
	slice1[3] = 2
	slice1[4] = 3

	return slice1
}

func main() {}
