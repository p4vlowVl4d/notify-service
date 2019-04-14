package handlers

import (
	json "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	config "github.com/spf13/viper"
)

func methodIndex(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var body indexRequest
	jsonErr := json.Unmarshal(bytes, &body)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	resp := MakeResponse(body)
	r.Header.Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(resp)
}

func MakeResponse(r indexRequest)*testResponse {
	a := config.GetString(fmt.Sprintf("routes.%s", r.MethodName))
	return &testResponse{
		Title: r.MethodName,
		About: a,
	}
}

type testResponse struct {
	Title	string `json:"title"`
	About	string `json:"about"`
}

type indexRequest struct {
	MethodName	string	`json:"title"`
}