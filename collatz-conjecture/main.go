package main

import (
	"errors"
	"fmt"
)

// Conjecture given a starting positive integer it returns the number of steps until 1 is reached
func Conjecture(n int) (steps int, err error) {
	steps = 0
	if n == 0 || n < 0 {
		err = errors.New("Need number greater than zero")
		return steps, err
	}
	for n > 1 {
		// increment steps
		steps++
		// check even
		if n%2 == 0 {
			n = n / 2
		} else {
			// odd
			n = (3 * n) + 1
		}
	}
	return steps, err
}

func main() {
	// input := 1
	// input2 := 16
	// input3 := 12
	input4 := 0
	fmt.Println(Conjecture(input4))
}
