package web

import (
	"encoding/json"
	"fmt"
	config "github.com/spf13/viper"
	"net/http"
)

func methodIndex(w http.ResponseWriter, r *http.Request) {
	var body indexRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
	}
	resp := MakeResponse(body)
	r.Header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func MakeResponse(r indexRequest)*testResponse {
	a := config.GetString(fmt.Sprintf("routes.%s.about", r.MethodName))
	if a == "" {
		return &testResponse{}
	}
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
	MethodName	string	`json:"method"`
}
