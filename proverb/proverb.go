package main

import (
	"fmt"
	"math/rand"
)

/* First Iteration
func Proverb(rhyme []string) []string {
	rand.Seed(42)
	// loop len of given array and output rand int
	// return formated string with two list items in it
	// return []string{}
	return fmt.Printf("For want of a %s the %s was lost.", rhyme[rand.Intn(len(rhyme))], rhyme[rand.Intn(len(rhyme))])
}
*/

// Second Iteration
// Solution idea from saurabhgarg-sg https://exercism.io/my/solutions/4ae3b779f9b2438793198ae1d04531f2
// Package proverb that displays pharases was given input
// Proverb generates the substituted rhyme
func Proverb(rhyme []string) []string {
	// array literal later used for your output
	base := [100]string{}
	inputLength := len(rhyme)
	// the output uses a slice you will later add rhymes too
	output := base[:0]

	// It returns empty string on empty input
	if inputLength == 0 {
		return output
	}
  // Given an array with input more than one input add to the output array slice and output to rhyme with given input
	if inputLength > 1 {
		for index := 0; index < inputLength-1; index++ {
			output = append(output, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[index], rhyme[index+1]))
		}
	}
  // Print and append final rhyme
	output = append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return output
}
