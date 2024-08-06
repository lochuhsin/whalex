package main

import (
	"net/http"
	"whalex/internal"

	log "github.com/sirupsen/logrus"
)

func register(mux *http.ServeMux) {
	mux.HandleFunc("/", internal.RootHandler)
}

func main() {
	mux := http.NewServeMux()
	internal.InitUpstreamRoute()
	register(mux)
	log.Info("Welcome to whalex, start routing")
	http.ListenAndServe(":8000", mux)
}
