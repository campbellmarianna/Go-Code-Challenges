/*
Bob is a lackadaisical teenager. In conversation, his responses are very limited.

Bob answers 'Sure.' if you ask him a question, such as "How are you?".

He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).

He answers 'Calm down, I know what I'm doing!' if you yell a question at him.

He says 'Fine. Be that way!' if you address him without actually saying anything.

He answers 'Whatever.' to anything else.

Bob's conversational partner is a purist when it comes to written communication and always follows normal rules regarding sentence punctuation in English.
*/
// Input: "YELL AT HIM (in all capitals)"
// Output: 'Whoa, chill out!'
package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
	"unicode"
)

// Hey function given a remark returns a lackadaisical response
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	} else if isAQuestion(remark) && isInAllCaps(remark) {
		return "Calm down, I know what I'm doing!"
	} else if isAQuestion(remark) {
		return "Sure."
	} else if isInAllCaps(remark) {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}

func isAQuestion(remark string) bool {
	// Ask a question
	lastChar := getLastRune(remark, 1)
	return lastChar == "?"
}

// Function Inspired by guitarrapc -  https://stackoverflow.com/questions/54475723/which-is-better-to-get-the-last-x-character-of-a-golang-string
func getLastRune(remark string, c int) string {     
	//DecodeLastRuneInString
	j := len(remark)
	for i := 0; i < c && j > 0; i++ {
		_, size := utf8.DecodeLastRuneInString(remark[:j])
		j -= size
	}
	lastByRune := remark[j:]
	// fmt.Println(lastByRune)
	return lastByRune
}

func isInAllCaps(remark string) bool {
	// Yell a remark
	// First make sure remark is all letters
	remark = strings.TrimLeftFunc(remark, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	if remark == "" {
		return false
	}
	firstChar := string(remark[0:1])
	UppercaseFirstCharacter := strings.ToUpper(firstChar)
	if firstChar == UppercaseFirstCharacter {
		numberOfLowercaseLetter := 0
		// check if the whole remark is uppercase
		for i := 0; i < len(remark); i++ {
				numberOfLowercaseLetter++
				// if we find a lowercase letter this is not a yelling statement
				currChar := string(remark[i])
				if (currChar != strings.ToUpper(currChar) && currChar == strings.ToLower(currChar)) {
					return false
				}
			}
		if numberOfLowercaseLetter == len(remark) {
			return true
		}
	}	
	return false
}


func main(){
	statement := "1, 2, 3"
	fmt.Println(Hey(statement))

}
