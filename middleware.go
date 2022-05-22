package main

import (
	"fmt"
	"net/http"
	"strings"
)

func checkTokenAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		var token string

		if authHeader != "" {
			if strings.Contains(authHeader, "Bearer") {
				bearerString := strings.Fields(authHeader)
				if len(bearerString) > 1 {
					token = bearerString[1]
				}
			} else {
				token = authHeader
				fmt.Println("Token 2", token)
			}
			//TODO: change hardcoded token
			if "Replace-This-Token" != token {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "No access token provided", http.StatusUnauthorized)
		}

	})
}
