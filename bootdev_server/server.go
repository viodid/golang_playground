package main

import (
	"errors"
	"net/http"
)

func server() {
	api := apiConfig{}

	mux := http.NewServeMux()

	handler := http.FileServer(http.Dir('.'))
	handler = http.StripPrefix("/app", handler)

	mux.Handle("/app/", api.middlewareMetricsInc(handler))
	mux.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		request.Header = map[string][]string{
			"Content-Type": {"text/plain", "charset=utf-8"},
		}
		writer.WriteHeader(200)
		_, err := writer.Write([]byte("OK"))
		if err != nil {
			return
		}
	})
	mux.HandleFunc("/reset", api.reset)

	mux.HandleFunc("/metrics", api.handler)

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
