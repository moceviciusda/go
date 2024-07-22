package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("squere does not exist on chessboard")
	}
	grains := math.Pow(2, float64(number-1))
	return uint64(grains), nil
}

func Total() (total uint64) {
	for i := 1; i < 65; i++ {
		g, _ := Square(i)
		total += g
	}

	return
}
