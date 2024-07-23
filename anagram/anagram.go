package anagram

import "strings"

func Detect(subject string, candidates []string) (anagrams []string) {
	subject = strings.ToLower(subject)

	for _, c := range candidates {
		if len(c) != len(subject) {
			continue
		}

		cLower := strings.ToLower(c)
		if subject == cLower {
			continue
		}

		cRunes := []rune(cLower)

		for _, r := range subject {
			for i, cr := range cRunes {
				if r != cr {
					continue
				}

				if len(cRunes) > i+1 {
					cRunes = append(cRunes[:i], cRunes[i+1:]...)
				} else {
					cRunes = cRunes[:i]
				}
			}
		}

		if len(cRunes) == 0 {
			anagrams = append(anagrams, c)
		}
	}

	return
}
