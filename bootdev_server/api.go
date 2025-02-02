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
	// TODO: middleware does not activate or doesn't update fileServerHits
	c.fileServerHits.Store(c.fileServerHits.Add(15))
	return next
}

func (c *apiConfig) handler(writer http.ResponseWriter, request *http.Request) {
	request.Header = map[string][]string{
		"Content-Type": {"text/plain", "charset=utf-8"},
	}
	body := fmt.Sprintf("Hits: %d\n", c.fileServerHits.Load())
	_, err := writer.Write([]byte(body))
	if err != nil {
		return
	}
}
