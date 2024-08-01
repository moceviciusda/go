package prime

func Factors(n int64) (factors []int64) {
	for factor := 2; n > 1; {

		if n%int64(factor) == 0 {
			factors = append(factors, int64(factor))
			n /= int64(factor)
			continue
		}

		factor++
	}

	return factors
}
