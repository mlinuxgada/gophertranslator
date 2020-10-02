package gophertranslator

import (
	"errors"
	"strings"
)

// TranslateWord - method for "translating" normal english word into gopher lang
// using several rules
func TranslateWord(word string) (string, error) {

	if len(word) == 0 {
		return word, errors.New("Given word is empty string")
	}

	gWord, match := ProcessWithPrefixG(word)

	if match {
		return gWord, nil
	}

	gWord, match = ProcessWithPrefixXR(word)

	if match {
		return gWord, nil
	}

	gWord, match = ProcessWithStartedConsonantAndPrefixQU(word)

	if match {
		return gWord, nil
	}

	// this must be last
	gWord, match = ProcessWithStartedConsonant(word)

	if match {
		return gWord, nil
	}

	return gWord, nil
} 

// ProcessWithPrefixG - checks if word starts with vowel letter
// and adds g prefix
func ProcessWithPrefixG(word string) (string, bool) {

	switch word[:1] {
    case "a", "e", "i", "o", "u", "y", "A", "E", "I", "O", "U", "Y":
		word = "g" + word
		return word, true
    }

	return word, false
}

// ProcessWithPrefixXR - checks if starts with xf and adds "ge" prefix
func ProcessWithPrefixXR(word string) (string, bool) {

	if strings.HasPrefix(word, "xr") {
		word = "ge" + word
		return word, true
	}

	return word, false
}

// ProcessWithStartedConsonant - check if word starts with consonant, check further
// and stop on the first vowel char
func ProcessWithStartedConsonant(word string) (string, bool) {

	pos := 0

	for i, ch := range word {

		switch string(ch) {

			case "a", "e", "i", "o", "u", "y", "A", "E", "I", "O", "U", "Y":

				if i == 0 {
					return word, false
				}

				pos = i - 1
				break
		}
	}

	gWord := word[pos:] + word[0:pos] + "ogo"

	return gWord, true
}

// ProcessWithStartedConsonantAndPrefixQU - check if starts with consonant
// AND has 2nd and 3rd chars are "qu"
func ProcessWithStartedConsonantAndPrefixQU(word string) (string, bool) {	

	switch string(word[0]) {

		case "a", "e", "i", "o", "u", "y", "A", "E", "I", "O", "U", "Y":
			return word, false
	}

	if len(word) < 3 {
		return word, false
	}

	if string(word[1:3]) != "qu" {

		return word, false
	}

	gWord := word[3:] + word[0:3] + "ogo"

	return gWord, true
}
