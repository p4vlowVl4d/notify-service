package web

import (
	"fmt"
	. "net/http"
)

func Router(w ResponseWriter, r *Request) {
	handle := indexHandle{
		name: "index",
	}
	handle.Run(w, r.Header)
}
type WebHandle interface {
	Run(w ResponseWriter, r Header)
}

type indexHandle struct {
	name string
}

func (i indexHandle) Run(w ResponseWriter, r Header) {
	fmt.Fprintf(w, "This is a %s for content-type: %s ", i.name, r.Get("content-type"))
}

type RequestData struct {
	Header  string
}
