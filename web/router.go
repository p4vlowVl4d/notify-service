package web

import (
	"fmt"
	. "net/http"
)

func Router(w ResponseWriter, r *Request) {
	handle := indexHandle{
		name: "index",
	}
	params := make(map[string]string)
	params["name"] = "index"
	data := RequestData{
		Method: "index",
		Params: params,
		Role:   "admin",
	}
	handle.Run(w, data)
}

type WebHandle interface {
	Run(w ResponseWriter, data RequestData)
}

type indexHandle struct {
	name string
}

func (i indexHandle) Run(w ResponseWriter, data RequestData) {
	fmt.Fprintf(w, "Method %s Role %s", data.Method, data.Role)
}

type RequestData struct {
	Method string
	Params map[string]string
	Role   string
}
