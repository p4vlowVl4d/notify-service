package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	config "src/github.com/spf13/viper"
)

const (
	defaultConfig = "/config/config.yaml"
	defaultPortListen = ":3000"
)

func main() {
	config.SetConfigFile(defaultConfig)
	if err := config.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	port := config.GetString("server.port")
	if len(port) > 0 {
		fmt.Printf("Starting server on %s", port)
		log.Fatal(http.ListenAndServe(port, nil))
	} else {
		fmt.Printf("Starting server on %s", defaultPortListen)
		log.Fatal(http.ListenAndServe(port, nil))
	}

	http.Handle("/", web.Route(addr))
}

type Router struct {
	name 	string
	handlers map[string]web.WebHandle
}