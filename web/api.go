package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/p4vlowVl4d/notify-service/web"
	"log"
	"net/http"
	"time"
)

func Execute(c *WebConf) {
	fmt.Println("Configure web server...")
	r := mux.NewRouter()
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         port,
		Handler:      r,
	}
	r.Get("/").Handler(web.ApiServe)
	fmt.Printf("Starting web server, listen on port: %s", port)
	log.Println(srv.ListenAndServe())
}

type WebConf struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PortListen   string
	Host         string
}
