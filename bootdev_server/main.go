package main

import (
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

}
