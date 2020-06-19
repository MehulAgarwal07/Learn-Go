package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/acha", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello WOrld"))
	})
	http.ListenAndServe("localhost:3000", mux)
}