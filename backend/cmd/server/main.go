package main

// Entry point for the Systems Monitor backend server
//
// Docs:
// - net/http: https://pkg.go.dev/net/http
// - chi router: https://pkg.go.dev/github.com/go-chi/chi/v5
// - CORS middleware: https://pkg.go.dev/github.com/go-chi/cors

import (
	"log"
	"net/http"
	"shelter/backend/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
	}))

	handlers.RegisterRoutes(r)

	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
