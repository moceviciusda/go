package reverse

import "strings"

func Reverse(input string) (reversed string) {
	chars := strings.Split(input, "")
	for i := len(chars) - 1; i > -1; i-- {
		reversed += chars[i]
	}

	return
}
