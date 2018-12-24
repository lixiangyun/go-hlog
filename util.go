package hlog

import (
	"fmt"

	"strings"
	"unicode"
)

func ignoreNumber(b rune) bool {
	return !(unicode.IsNumber(b) && unicode.IsSpace(b))
}

// 10 KB MB GB
func parseSize(line string) (int, error) {
	var number int

	cnt, err := fmt.Sscanf(line, "%d", &number)
	if cnt != 1 {
		return 0, err
	}

	idx := strings.LastIndexFunc(line, ignoreNumber)

	if 0 == strings.Compare(line[idx:], "KB") {
		number *= 1024
	} else if 0 == strings.Compare(line[idx:], "MB") {
		number *= 1024 * 1024
	} else if 0 == strings.Compare(line[idx:], "GB") {
		number *= 1024 * 1024 * 1024
	}

	return number, nil
}
