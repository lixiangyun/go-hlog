package hlog

import (
	"errors"
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

/* %-20.30c */
type Width struct {
	left bool /* true */
	max  int  /* 30 */
	min  int  /* 20 */
}

/*
 * parm
 *  string: %-20.30c
 *
 * ret
 *  int   :parse width spec offset
 *  width :
 */
func parseWidth(line string, w *Width) (int, error) {
	var begin int
	var end int

	if len(line) < 2 || line[0] != '%' {
		return 0, nil
	}

	if line[1] == '-' {
		w.left = true
		end = 2
	} else {
		end = 1
	}

	begin = end
	for _, v := range line[begin:] {
		if !unicode.IsNumber(v) {
			break
		}
		end++
	}

	if begin < end {

		cnt, err := fmt.Sscanf(line[begin:end], "%d", &w.min)
		if err != nil {
			return 0, err
		}

		if cnt != 1 {
			w.min = 0
		}
	}

	if line[end] == '.' {

		end++
		begin = end
		for _, v := range line[begin:] {
			if !unicode.IsNumber(v) {
				break
			}
			end++
		}

		if begin < end {
			cnt, err := fmt.Sscanf(line[begin:end], "%d", &w.max)
			if err != nil {
				return 0, err
			}

			if cnt != 1 {
				w.max = 0
			}
		}
	}

	if len(line) < end || !unicode.IsLetter(rune(line[end+1])) {
		return 0, errors.New(fmt.Sprintf("parse width[%d] failed!", line[:end]))
	}

	return end, nil
}
