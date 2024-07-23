package proverb

import "fmt"

func Proverb(rhyme []string) (res []string) {
	for i, word := range rhyme {
		if i >= len(rhyme)-1 {
			res = append(res, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
		} else {
			res = append(res, fmt.Sprintf("For want of a %v the %v was lost.", word, rhyme[i+1]))
		}
	}

	return
}
