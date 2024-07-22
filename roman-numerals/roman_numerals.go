package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("invalid number")
	}

	var res string

	thousands := input / 1000
	input %= 1000

	for i := 0; i < thousands; i++ {
		res += "M"
	}

	if input < 900 {

		halfThousands := input / 500
		input %= 500
		for i := 0; i < halfThousands; i++ {
			res += "D"
		}

		hundreds := input / 100
		input %= 100
		if hundreds > 3 {
			res += "CD"
		} else {
			for i := 0; i < hundreds; i++ {
				res += "C"
			}
		}
	} else {
		res += "CM"
		input -= 900
	}

	if input < 90 {
		fifties := input / 50
		input %= 50
		for i := 0; i < fifties; i++ {
			res += "L"
		}

		tens := input / 10
		input %= 10
		if tens > 3 {
			res += "XL"
		} else {
			for i := 0; i < tens; i++ {
				res += "X"
			}
		}
	} else {
		res += "XC"
		input -= 90
	}

	if input < 9 {
		fives := input / 5
		input %= 5
		for i := 0; i < fives; i++ {
			res += "V"
		}

		if input > 3 {
			res += "IV"
		} else {
			for i := 0; i < input; i++ {
				res += "I"
			}
		}
	} else {
		res += "IX"
	}

	return res, nil
}
