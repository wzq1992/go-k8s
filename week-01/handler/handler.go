package handler

import (
	"fmt"
	"go-k8s/pkg/env"
	"log"
	"net/http"
)

func HeadersHandler(w http.ResponseWriter, r *http.Request) {
	count := 0
	for key, vals := range r.Header {
		if len(vals) > 0 {
			count++
			w.Header().Set(key, vals[0])
		}
	}
	writeString(w, fmt.Sprintf("write headers: [%d]", count))
	logRequestInfo(r, 200)
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	code := 200
	version := env.GetVersion()
	if version != "" {
		w.Header().Set("X-Version", env.GetVersion())
		writeString(w, "version: "+version)
	} else {
		code = 204
		writeString(w, "version not set")
		w.WriteHeader(code)
	}

	logRequestInfo(r, code)
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	code := 200
	writeString(w, "ok")
	logRequestInfo(r, code)
	w.WriteHeader(code)
}

func logRequestInfo(r *http.Request, code int) {
	log.Printf("Request Url: %s, Client IP: %s, Code: %d", r.URL.String(), r.RemoteAddr, code)
}

func writeString(w http.ResponseWriter, data string) {
	_, _ = w.Write([]byte(data))
}
