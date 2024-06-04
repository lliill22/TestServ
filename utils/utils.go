package utils

import "net/http"

func GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Welcome GET")) }
}

func PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome POST"))
	}
}
