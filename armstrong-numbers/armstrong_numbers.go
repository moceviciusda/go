package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	var powSum float64
	sn := strconv.Itoa(n)
	for _, r := range sn {
		powSum += math.Pow(float64(r-'0'), float64(len(sn)))
	}

	return powSum == float64(n)
}
