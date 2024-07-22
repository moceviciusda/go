package isbn

import (
	"strings"
	"unicode/utf8"
)

func IsValidISBN(isbn string) bool {
	clean := strings.ReplaceAll(isbn, "-", "")
	if len(clean) != 10 {
		return false
	}

	checkChar, _ := utf8.DecodeLastRuneInString(clean)

	var checkNum int
	if checkChar == 'X' {
		checkNum = 10
	} else {
		checkNum = int(checkChar - '0')
	}

	var res int
	multiplier := 10
	for i, r := range clean {
		if i != len(clean)-1 {
			res += multiplier * int(r-'0')
		} else {
			res += checkNum
		}

		multiplier--
	}

	return res%11 == 0

}
