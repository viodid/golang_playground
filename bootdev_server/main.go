package main

import (
	"errors"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler := http.FileServer(http.Dir('.'))
	handler = http.StripPrefix("/app", handler)
	mux.Handle("/app/", handler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		if !errors.Is(http.ErrServerClosed, err) {
			panic(err)
		}
	}
}
