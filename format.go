package hlog

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

type TRANSFER func(*buffer, *spec, *format, *event) error

type spec struct {
	name string
	body string
	max  int
	min  int /**/

	fun TRANSFER
}

type format struct {
	name string
	list []spec
}

type formats_map struct {
	sync.RWMutex
	list map[string]format
}

var g_hlogHostName string
var g_hlogPid string

func init() {
	g_hlogPid = fmt.Sprintf("%d", os.Getpid())
	g_hlogHostName, _ = os.Hostname()
}

/* %d() */
func SpecTimeStamp(buf *buffer, sp *spec, f *format, evt *event) error {
	buf.append(evt.timestamp.Format("RFC1123"))
	return nil
}

func SpecPercent(buf *buffer, sp *spec, f *format, evt *event) error {
	buf.append("%")
	return nil
}

func SpecEnd(buf *buffer, sp *spec, f *format, evt *event) error {
	buf.append("\r\n")
	return nil
}

func SpecUsermsg(buf *buffer, sp *spec, f *format, evt *event) error {
	buf.append(evt.context)
	return nil
}

func SpecHostname(buf *buffer, sp *spec, f *format, evt *event) error {
	buf.append(g_hlogHostName)
	return nil
}

func SpecPid(buf *buffer, sp *spec, f *format, evt *event) error {
	buf.append(g_hlogPid)
	return nil
}

func SpecLine(buf *buffer, sp *spec, f *format, evt *event) error {
	_, _, line, b := runtime.Caller(3)
	if b == true {
		buf.append(fmt.Sprintf("line:%v", line))
	}
	return nil
}

func SpecFile(buf *buffer, sp *spec, f *format, evt *event) error {
	_, file, _, b := runtime.Caller(3)
	if b == true {
		buf.append(fmt.Sprintf("file:%v", file))
	}
	return nil
}

func SpecFunction(buf *buffer, sp *spec, f *format, evt *event) error {
	ptr, _, _, b := runtime.Caller(3)
	if b == true {
		buf.append(fmt.Sprintf("function:%v", ptr))
	}
	return nil
}

func SpecClass(buf *buffer, sp *spec, f *format, evt *event) error {
	buf.append(evt.class)
	return nil
}

func SpecStack(buf *buffer, sp *spec, f *format, evt *event) error {
	tmpbuff := make([]byte, 1024)
	cnt := runtime.Stack(tmpbuff, false)
	buf.append(fmt.Sprintf("callstack:%d\r\n%s\r\n", cnt, string(tmpbuff)))
	return nil
}

func SpecString(buf *buffer, sp *spec, f *format, evt *event) error {
	buf.append(sp.body)
	return nil
}

func SpecLevelLowercase(buf *buffer, sp *spec, f *format, evt *event) error {
	lv := LevelGetById(evt.level)
	if lv.id == 0 {
		buf.append("*")
	} else {
		buf.append(lv.name_lower)
	}
	return nil
}

func SpecLevelUppercase(buf *buffer, sp *spec, f *format, evt *event) error {
	lv := LevelGetById(evt.level)
	if lv.id == 0 {
		buf.append("*")
	} else {
		buf.append(lv.name_upper)
	}
	return nil
}

func NewFormat(name string, fmat string) error {

	f := new(format)
	f.name = name

	return nil
}
