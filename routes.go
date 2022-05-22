package main

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	assetsFs := http.FileServer(http.FS(assets))

	router := http.NewServeMux()
	router.Handle("/assets/", assetsFs)
	router.HandleFunc("/api/health", health)
	colorHandler := http.HandlerFunc(color)
	router.Handle("/api/v1/color", checkTokenAuth(colorHandler))
	router.HandleFunc("/readme.html", readme)
	router.HandleFunc("/", index)

	return router
}
