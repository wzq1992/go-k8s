package main

import (
	"fmt"
	"go-k8s/handler"
	"net/http"
)

func main() {
	fmt.Println("http server listen at: http://127.0.0.1:8080")

	http.HandleFunc("/headers", handler.HeadersHandler)
	http.HandleFunc("/version", handler.VersionHandler)
	http.HandleFunc("/healthz", handler.HealthzHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("server start failed.")
	}
}
