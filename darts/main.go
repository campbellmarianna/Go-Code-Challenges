package main

// inputs:
// outputs:

import (
	"fmt"
	"math"
)

// Score returns the points earned based on the given position of a tossed dart
func Score(x float64, y float64) int {
	// create score
	var score int
	// figure out the radius give x and y
	dl := math.Pow(x, 2) + math.Pow(y, 2)
	dl = math.Sqrt(dl)
	fmt.Println("Radius:", dl)
	// figure out where dart landed and apply associated score
	if dl > 10 {
		// outside circle
		// no points earned
	}
	if dl <= 10 && dl > 5 {
		// in outside circle
		score++
	}
	if dl <= 5 && dl > 1 {
		// middle circle
		score += 5
	}
	if dl <= 1 {
		// inner circle
		score += 10
	}
	return score
}

func main() {
	// dartTossX, dartTossY := -9.0, 9.0
	// dartTossX, dartTossY := 0.0, 10.0
	// dartTossX, dartTossY := -5.0, 0.0
	// dartTossX, dartTossY := 0.0, 0.0
	// dartTossX, dartTossY := 0.8, -0.8
	// dartTossX, dartTossY := -3.5, 3.5
	// dartTossX, dartTossY := -3.6, 3.6
	dartTossX, dartTossY := 0.0, -1.0
	fmt.Println(Score(dartTossX, dartTossY))
}
