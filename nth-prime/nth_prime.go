package prime

import "errors"

func IsPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return -1, errors.New("'n' is less than 1")
	}

	var pc int

	for i := 2; pc < n; i++ {
		if IsPrime(i) {
			pc++
		}

		if pc >= n {
			return i, nil
		}
	}

	return -1, errors.New("unexpected error")
}
