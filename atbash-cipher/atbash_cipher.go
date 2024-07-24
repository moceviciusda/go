package atbash

import (
	"unicode"
)

// const (
// 	alphabet = abcdefghijklmnopqrstuvwxyz
// 	tebahpla = zyxwvutsrqponmlkjihgfedcba
// )

func Atbash(s string) string {
	var runes []rune
	ci := 0
	for _, r := range s {
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			continue
		}

		if ci != 0 && ci%5 == 0 {
			runes = append(runes, ' ')
		}

		if unicode.IsDigit(r) {
			runes = append(runes, r)
			ci++
		}
		if unicode.IsLetter(r) {
			runes = append(runes, 'z'-(unicode.ToLower(r)-'a'))
			ci++
		}
	}

	return string(runes)
}
