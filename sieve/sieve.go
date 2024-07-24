package sieve

func Sieve(limit int) (primes []int) {
	marked := make([]bool, limit+1)

	for i := 2; i <= limit; i++ {
		if marked[i] {
			continue
		}

		primes = append(primes, i)

		for j := 2; j*i <= limit; j++ {
			marked[j*i] = true
		}
	}

	return
}
