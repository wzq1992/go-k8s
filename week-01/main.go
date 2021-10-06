package main

import (
	"go-k8s/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/headers", handler.HeadersHandler)
	http.HandleFunc("/version", handler.VersionHandler)
	http.HandleFunc("/healthz", handler.HealthzHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("server start failed.")
	}
}
