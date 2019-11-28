package main

import (
	"fmt"
)

// ToRNA returns the correspnding RNA to a given DNA strand
func ToRNA(dna string) string {
	// check empty return empty
	if dna == "" {
		return dna
	} 
	var rna string
	// store a map with key curr_dna to value rna
	var DNAs = map[string]string{
		"C":"G",
		"G":"C",
		"T":"A",
		"A":"U",
	}

	var dnaPart string

	for _,c := range dna {
		dnaPart = string(c)
		// when I find curr_dna in map
		rnaValue := DNAs[dnaPart]
		// add its value to a rna string
		rna += rnaValue
	}
		
	return rna
}

func main(){
	dnaStrand := "C"
	fmt.Println(ToRNA(dnaStrand))
}