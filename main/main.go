package main

import (
	"TestServ/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/api", utils.GetHandler())
	router.Post("/api", utils.PostHandler())
	http.ListenAndServe(":9000", router)
}
