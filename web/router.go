package web

import (
	"net/http"
	"src/github.com/spf13/viper"
)

func Route(addr string) http.Handler {

}

func getHandlers(routes []string) map[string]WebHandle {

}

type WebHandle struct {
	name string
	authEnable bool
	worker	http.Handler
}

