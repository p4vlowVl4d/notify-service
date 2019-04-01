package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/p4vlowVl4d/notify-service/web"
	config "github.com/spf13/viper"
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
	http.HandleFunc("/", web.Router)
	port := config.GetString("server.port")
	if len(port) > 0 {
		fmt.Printf("Starting server on %s\n", port)
		log.Fatal(http.ListenAndServe(port, nil))
	} else {
		fmt.Printf("Starting server on %s\n", defaultPortListen)
		log.Fatal(http.ListenAndServe(port, nil))
	}

}

//type Router struct {
//	name 	string
//	handlers map[string]web.WebHandle
//}