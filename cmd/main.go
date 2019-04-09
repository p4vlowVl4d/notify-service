package main

import (
	"fmt"
	"github.com/p4vlowVl4d/notify-service/web"
	config "github.com/spf13/viper"
	"log"
	"net/http"
)

const (
	defaultConfig = "config/config.yml"
	defaultPortListen = ":3000"
)

func main() {
	config.SetConfigFile(defaultConfig)
	if err := config.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	port := config.GetString("server.port")
	fmt.Printf("Starting web server, listen on port: %s", port)
	http.ListenAndServe(port, web.requestHandle)
}

type WebData struct {

}