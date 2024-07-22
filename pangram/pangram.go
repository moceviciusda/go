package pangram

import "unicode"

const lettersInAlphabet = 26

func IsPangram(input string) bool {
	letters := map[rune]int{}

	for _, r := range input {
		rCap := unicode.ToLower(r)
		if unicode.IsLetter(rCap) {
			letters[rCap]++
		}
	}

	return len(letters) == lettersInAlphabet
}
