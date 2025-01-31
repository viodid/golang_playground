package main

import (
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileServerHits atomic.Int32
}

func (c *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	// TODO: middleware does not activate or doesn't update fileServerHits
	c.fileServerHits.Add(1)
	return next
}

func (c *apiConfig) handler(writer http.ResponseWriter, request *http.Request) {
	request.Header = map[string][]string{
		"Content-Type": {"text/plain", "charset=utf-8"},
	}
	body := "Hits: " + string(byte(c.fileServerHits.Load()+48))
	_, err := writer.Write([]byte(body))
	if err != nil {
		return
	}
}
