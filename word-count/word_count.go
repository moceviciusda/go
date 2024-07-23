package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := Frequency{}

	re := regexp.MustCompile(`\w+('\w+)*`)
	words := re.FindAllString(phrase, -1)

	for _, word := range words {
		freq[strings.ToLower(word)]++
	}

	return freq
}
