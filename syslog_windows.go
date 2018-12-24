package hlog

import (
	"errors"
)

func syslogAtoi(string) (int, error) {
	return 0, errors.New("windows not support syslog!")
}
