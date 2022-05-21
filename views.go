package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write(indexPage)
}

func readme(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write(readmePage)
}

func color(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(ws2812Color)
	case "POST":
		var req Color

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(req)
		ws2812Color = req
		json.NewEncoder(w).Encode(ws2812Color)
	}
}
