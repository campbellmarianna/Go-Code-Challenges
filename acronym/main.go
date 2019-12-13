// Package main provides an abbreviation for a given phrase
package main

import(
	"fmt"
	"strings"
	"unicode"
)
// input: "Portable" "Network" "Graphics"
// output: PNG

// Abbreviate given a phrase returns an acronym
func Abbreviate(s string) string {
	// split the given string into a slices of items separated by spaces or dashes
	r := []rune("'")
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && c != r[0]
	}
	words := strings.FieldsFunc(s,f)
	// loop through each word, captilize it and store it in the result 
	var acronym string
	for _,word := range words {
		firstLetter := word[0:1]
		firstLetter = strings.ToUpper(firstLetter)
		acronym += firstLetter

	}
	return acronym
}

func main(){
	// phrase := "Portable Network Graphics"
	// phrase2 := "Ruby on Rails"
	phrase3 := "Halley's Comet"
	fmt.Println(Abbreviate(phrase3))
}