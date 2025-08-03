package main

import (
	"log"
	"net/http"
)

func (app *Application) enableCORS(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token, Authorization")
			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (app *Application) requestLogger(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("received request: from[%s] - [%s] %s %s \n", r.RemoteAddr, r.Method, r.Host, r.URL.Path)

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
