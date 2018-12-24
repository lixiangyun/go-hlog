package hlog

import (
	"time"
)

type event struct {
	file      string
	line      int
	statck    string
	context   string
	timestamp time.Time
}
