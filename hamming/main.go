package main

import (
	"fmt"
)

// Distance returns the Hamming of two DNA
func Distance(a, b string) (int, error) {
	// identifylengths of DNA a and b are unequals
	if len(a) != len(b) {
		return 0, fmt.Errorf("The lengths of DNA are different - a: %d, b %d", len(a), len(b))
	}
	// return how many differences in DNA there are	
	cnt := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt++
		}
	}
	return cnt, nil
}

func main() {
	var a = "ABCCB"
	var b = "BBCAB"
	fmt.Println(Distance(a, b))
}
