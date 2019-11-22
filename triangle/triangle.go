// Package main implements a simple check for if a triangle is equilateral, isosceles, scalene
package main

import (
	"fmt"
	"math"
)

type Kind string

const (
    NaT = "Nat"     // not a triangle
	  Equ = "Equ"     // equilateral
	  Iso = "Iso"     // isosceles
	  Sca = "Sca"
)

// KindFromSides returns Equ for equilateral, Iso for isosceles or Sca scalene based on the given triangle sides
// func KindFromSides(a, b, c float64) Kind {
// 	var k Kind
// 	// check for invalid triangle
// 	if a + b <= c {
// 		k = Kind{"NaT"}
// 		return k
// 	}
// 	counter := 0
// 	sides := [3]float64{a, b, c}
// 	// create map
// 	kvs := map[float64]float64{a:0, b:0, c:0}
// 	for _, side := range sides {
// 		kvs[side]++
// 	} 

// 	for _,side := range sides {
// 		// check Equ
// 		if kvs[side] == 3 {
// 			k = Kind{"Equ"}
// 		}
// 		// check Iso
// 		if kvs[side] == 2 {
// 			k = Kind{"Iso"}
// 		}
// 		// check Sca
// 		if kvs[side] == 1 {
// 			counter++
// 			if counter == 3 {
// 				k = Kind{"Sca"}
// 			} 
// 		}
// 	}

// 	return k
// }

// Iteration 2 Solution Idea Inspired by: https://exercism.io/tracks/go/exercises/triangle/solutions/3161f9c56b074c1b9e591695a3e4b484

// create a function to check every case for if a function is not valid

func isTriangleValid(a, b, c float64) bool {
	// a is negative or 0 or NaN or Infinity
	if a <= 0 || math.IsNaN(a) == true || math.IsInf(a, int(math.J1(a))) == true {
		return false
	}
		
	// b is negative or 0 or NaN or Infinity
	if b <= 0 || math.IsNaN(b) == true || math.IsInf(b, int(math.J1(b))) == true {
		return false
	}
		
	// c is negative or 0 or NaN or Infinity
	if c <= 0 || math.IsNaN(c) == true || math.IsInf(c, int(math.J1(c))) == true {
		return false
	}
		
	// a + b < c
	if a + b < c {
		return false
	}
		
	// b + c < a
	if b + c < a {
		return false
	}
		
	// c + a < b
	if c + a < b {
		return false
	}
	return true
}

// Once we know the triangle is valid we can check - 12 minutes
	// if all three sides are the same length
		// equalatorial
	// if two sides are the same
		// isosalces

func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if isTriangleValid(a, b, c) == false {
		k = NaT
		return k
	} 
	// check Equ
	if a == b && b == c && a == c {
		k = Equ
		return k
	}
	// check Iso
	if a == b || b == c || a == c {
		k = Iso
		return k
	}
	// otherwise it is scalene
	k = Sca
	return k
}

func main() {
	var a float64 = 3
	var b float64 = 4
	var c float64 = math.NaN()
	fmt.Println("C value:", c)
	fmt.Println(KindFromSides(a,b,c))
}