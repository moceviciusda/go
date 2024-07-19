package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) < 2 {
		return false
	}

	var sum int
	double := false
	for i := len(id) - 1; i >= 0; i-- {

		digit, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit

		double = !double
	}

	return sum%10 == 0
}
