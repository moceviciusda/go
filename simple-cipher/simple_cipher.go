package cipher

import (
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift int

type vigenere string

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}
	return shift(distance)
}

func (c shift) Encode(input string) (encoded string) {
	input = strings.ToLower(input)

	for _, r := range input {
		if unicode.IsLetter(r) {
			encoded += string('a' + (r-'a'+rune(c)+26)%26)
		}
	}

	return
}

func (c shift) Decode(input string) (decoded string) {
	input = strings.ToLower(input)

	for _, r := range input {
		if unicode.IsLetter(r) {
			decoded += string('a' + (r-'a'-rune(c)+26)%26)
		}
	}
	return
}

func NewVigenere(key string) Cipher {
	invalid := true
	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}

		if r > 'a' {
			invalid = false
		}

	}

	if invalid {
		return nil
	}

	return vigenere(key)
}

func (v vigenere) Encode(input string) (encoded string) {
	var sft shift

	li := 0
	for _, ir := range input {
		if unicode.IsLetter(ir) {
			sft = shift(v[li%len(v)] - 'a')

			encoded += sft.Encode(string(ir))

			li++
		}
	}

	return
}

func (v vigenere) Decode(input string) (decoded string) {
	var sft shift

	li := 0
	for _, ir := range input {
		if unicode.IsLetter(ir) {
			sft = shift(v[li%len(v)] - 'a')

			decoded += sft.Decode(string(ir))

			li++
		}
	}

	return
}
