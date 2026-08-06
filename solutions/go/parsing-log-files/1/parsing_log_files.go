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
	re := regexp.MustCompile(`<(\*|~|-|=)*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return string(re.ReplaceAll([]byte(text), []byte("")))
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+\w+`)
	reUserName := regexp.MustCompile(`User|\s+`)
	for i, line := range lines {
		if user_chunk := re.FindString(line); user_chunk != "" {
			userName := reUserName.ReplaceAllString(user_chunk, "")
			lines[i] = fmt.Sprintf("[USR] %s %s", userName, line)
		}
	}
	return lines
}
