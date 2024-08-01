package phonenumber

import (
	"fmt"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	var clean string
	for _, r := range phoneNumber {
		if unicode.IsDigit(r) {
			clean += string(r)
		}
	}

	if len(clean) > 11 || len(clean) < 10 {
		return clean, fmt.Errorf("number must be 11 or 10 digits long, len: %v", len(clean))
	}

	if len(clean) == 11 {
		if clean[0] != '1' {
			return clean, fmt.Errorf("invalid country code: %v", clean[0])
		}

		clean = clean[1:]
	}

	if int(clean[0]-'0') < 2 || int(clean[3]-'0') < 2 {
		return clean, fmt.Errorf("invalid number: %v", clean)
	}

	return clean, nil
}

func AreaCode(phoneNumber string) (string, error) {
	cleanNumber, err := Number(phoneNumber)
	if err != nil {
		return cleanNumber, err
	}

	return cleanNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {
	cleanNumber, err := Number(phoneNumber)
	if err != nil {
		return cleanNumber, err
	}

	return fmt.Sprintf("(%v) %v-%v", cleanNumber[:3], cleanNumber[3:6], cleanNumber[6:]), nil
}
