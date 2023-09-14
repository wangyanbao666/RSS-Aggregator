package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error parsing %v", payload)
		w.WriteHeader(code)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(dat)
	w.WriteHeader(code)
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error: ", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	responseJson(w, code, &errResponse{Error: msg})
}
