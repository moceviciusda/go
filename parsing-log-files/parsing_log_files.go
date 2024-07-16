package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~\*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (total int) {
	re := regexp.MustCompile(`".*(?i)password(?-i).*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			total++
		}
	}
	return
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)

	for i, line := range lines {
		user := re.FindStringSubmatch(line)
		if len(user) == 0 {
			continue
		}

		lines[i] = fmt.Sprintf("[USR] %s %s", user[1], lines[i])
	}

	return lines
}
