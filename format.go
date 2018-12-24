package hlog

type TRANSFER func(*buffer, *spec, *format, *event)

type spec struct {
	name string
	body string
	fun  TRANSFER
}

type format struct {
	name string
	list []spec
}

type formats struct {
	list map[string]format
}
