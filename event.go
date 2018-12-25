package hlog

import (
	"time"
)

type event struct {
	class     string
	file      string
	line      int
	function  string
	statck    string
	context   string
	level     int
	timestamp time.Time
}
