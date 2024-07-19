package isogram

import "unicode"

func IsIsogram(word string) bool {
	for i, letter := range word {
		if !unicode.IsLetter(letter) {
			continue
		}

		for _, l := range word[i+1:] {
			if unicode.ToLower(letter) == unicode.ToLower(l) {
				return false
			}
		}
	}

	return true
}
