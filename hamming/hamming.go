package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("string length is not equal")
	}

	var distance int
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
