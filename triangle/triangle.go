// Package triangle implements a simple check for if a triangle is equilateral, isosceles, scalene
package main

import (
	"fmt"
)
// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind struct {
	k string
}

// const (
    // Pick values for the following identifiers used by the test program.
    // NaT // not a triangle
    // Equ // equilateral
    // Iso // isosceles
    // Sca // scalene
// )

// KindFromSides returns Equ for equilateral, Iso for isosceles or Sca scalene based on the given triangle sides
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	// check for invalid triangle
	if a + b <= c {
		k = Kind{"NaT"}
		return k
	}
	counter := 0
	sides := [3]float64{a, b, c}
	// create map
	kvs := map[float64]float64{a:0, b:0, c:0}
	for _, side := range sides {
		kvs[side]++
	} 

	for _,side := range sides {
		// check Equ
		if kvs[side] == 3 {
			k = Kind{"Equ"}
		}
		// check Iso
		if kvs[side] == 2 {
			k = Kind{"Iso"}
		}
		// check Sca
		if kvs[side] == 1 {
			counter++
			if counter == 3 {
				k = Kind{"Sca"}
			} 
		}
	}

	return k
}

func main() {
	var a float64 = 3
	var b float64 = 4
	var c float64 = 5
	fmt.Println(KindFromSides(a,b,c))
}