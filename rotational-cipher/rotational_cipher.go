package rotationalcipher

func RotationalCipher(plain string, shiftKey int) (res string) {
	for _, r := range plain {
		if r >= 'a' && r <= 'z' {
			c := r + rune(shiftKey%26)
			if c > 'z' {
				c = 'a' + c - 'z' - 1
			}
			res += string(c)

		} else if r >= 'A' && r <= 'Z' {
			c := r + rune(shiftKey%26)
			if c > 'Z' {
				c = 'A' + c - 'Z' - 1
			}
			res += string(c)

		} else {
			res += string(r)
		}
	}

	return
}
