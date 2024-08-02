package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) (acronym string) {
	words := strings.Split(s, " ")
	for i, w := range words {
		sw := strings.Split(w, "-")
		if len(sw) > 1 {
			words = append(words[:i], append(sw, words[i+1:]...)...)
		}
	}

	for _, word := range words {
		for _, r := range word {
			if unicode.IsLetter(r) {
				acronym += string(unicode.ToUpper(r))
				break
			}
		}
	}

	return
}
