package main

import "net/http"

func handlerFunction(w http.ResponseWriter, r *http.Request) {
	responseJson(w, 200, struct{}{})
}
