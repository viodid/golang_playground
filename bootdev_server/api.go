package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileServerHits atomic.Int32
}

func (c *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.fileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (c *apiConfig) handler(w http.ResponseWriter, r *http.Request) {
	r.Header = map[string][]string{
		"Content-Type": {"text/plain", "charset=utf-8"},
	}
	body := fmt.Sprintf("Hits: %d\n", c.fileServerHits.Load())
	_, err := w.Write([]byte(body))
	if err != nil {
		return
	}
}

func (c *apiConfig) reset(writer http.ResponseWriter, request *http.Request) {
	c.fileServerHits.Store(0)
}
