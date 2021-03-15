package main

import (
	"http-proxy/proxy"
	"log"

	"net/http"
)

func main() {
	p := &proxy.Proxy{}

	if err := http.ListenAndServe("127.0.0.1:8080", p); err != nil {
		log.Fatalf(err.Error())
	}
}
