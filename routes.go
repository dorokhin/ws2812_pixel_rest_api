package main

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	assetsFs := http.FileServer(http.FS(assets))

	router := http.NewServeMux()
	router.Handle("/assets/", assetsFs)
	router.HandleFunc("/api/health", health)
	router.HandleFunc("/api/v1/color", color)
	router.HandleFunc("/readme.html", readme)
	router.HandleFunc("/", index)

	return router
}
