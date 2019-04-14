package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Execute(c *WebConf) {
	fmt.Println("Configure web server...")
	r := mux.NewRouter()
	r.Host(c.Host)
	SetHandlers(r)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf("%s:%s", c.Host, c.PortListen),
		Handler:      r,
	}
	fmt.Printf("Starting web server, listen on address: %s:%s", c.Host, c.PortListen)
	log.Println(srv.ListenAndServe())
}

func SetHandlers(r *mux.Router) {
	r.HandleFunc("/", methodIndex).
		Methods("GET").
		Headers("Content-Type", "application/json")
}

type WebConf struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PortListen   string
	Host         string
}

type WebRoute struct {
	Name         string
	Method       string
	Url          string
	ContentType  string
	Private, Api bool
}