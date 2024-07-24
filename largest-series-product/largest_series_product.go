package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) {
		return -1, errors.New("span must be smaller than string length")
	}
	if span < 1 {
		return -1, errors.New("span must be > 0")
	}

	var max int64
	for i := 0; i <= len(digits)-span; i++ {
		var product int64 = 1
		for _, digit := range digits[i : i+span] {
			di, err := strconv.Atoi(string(digit))
			if err != nil {
				return -1, errors.New("invalid digit")
			}

			product *= int64(di)
		}

		if product > max {
			max = product
		}
	}

	return max, nil
}
