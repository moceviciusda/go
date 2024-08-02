package twelve

import (
	"fmt"
	"strings"
)

const base = "On the %s day of Christmas my true love gave to me: %s"

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = []string{
	"twelve Drummers Drumming,",
	"eleven Pipers Piping,",
	"ten Lords-a-Leaping,",
	"nine Ladies Dancing,",
	"eight Maids-a-Milking,",
	"seven Swans-a-Swimming,",
	"six Geese-a-Laying,",
	"five Gold Rings,",
	"four Calling Birds,",
	"three French Hens,",
	"two Turtle Doves, and",
	"a Partridge in a Pear Tree.",
}

func Song() string {
	var song []string = make([]string, 12)
	for i := range days {
		song[i] = Verse(i + 1)
	}
	return strings.Join(song, "\n")
}

func Verse(day int) string {
	if (day < 1) || (day > 12) {
		return ""
	}
	return fmt.Sprintf(base, days[day-1], listOfGifts(day))
}

func listOfGifts(day int) string {
	return strings.Join(gifts[12-day:], " ")
}
