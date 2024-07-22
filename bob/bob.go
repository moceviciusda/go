package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")
	isUpperCase := strings.ToUpper(remark) == remark && strings.IndexFunc(remark, unicode.IsLetter) != -1

	if isUpperCase {
		if strings.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}

	if strings.HasSuffix(remark, "?") {
		return "Sure."
	}

	if remark == "" {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
