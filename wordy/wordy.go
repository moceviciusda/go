package wordy

import (
	"strconv"
	"strings"
)

const (
	Addition       = "plus"
	Subtraction    = "minus"
	Multiplication = "multiplied"
	Division       = "divided"
)

func Answer(question string) (result int, ok bool) {
	question = question[:len(question)-1]
	words := strings.Split(question, " ")

	for i := 0; i < len(words); i++ {
		switch strings.ToLower(words[i]) {
		case "what":
			if i != 0 {
				return
			}
		case "is":
			if i == len(words)-1 {
				return
			}
		case Addition:
			i++
			if i >= len(words) {
				return
			}

			num, err := strconv.Atoi(words[i])
			if err != nil {
				return
			}
			result += num

		case Subtraction:
			i++
			if i >= len(words) {
				return
			}

			num, err := strconv.Atoi(words[i])
			if err != nil {
				return
			}
			result -= num

		case Multiplication:
			i++
			if i >= len(words) || strings.ToLower(words[i]) != "by" {
				return
			}

			i++
			if i >= len(words) {
				return
			}

			num, err := strconv.Atoi(words[i])
			if err != nil {
				return
			}

			result *= num

		case Division:
			i++
			if i >= len(words) || strings.ToLower(words[i]) != "by" {
				return
			}

			i++
			if i >= len(words) {
				return
			}

			num, err := strconv.Atoi(words[i])
			if err != nil {
				return
			}

			result /= num

		default:
			num, err := strconv.Atoi(words[i])
			if err != nil {
				return
			}

			if strings.ToLower(words[i-1]) != "is" {
				return
			}

			result += num
		}
	}

	return result, true
}
