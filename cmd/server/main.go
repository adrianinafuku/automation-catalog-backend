package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"automation-catalog-backend/internal/api"
)

func main() {
	// Load .env if running locally
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()

	api.RegisterRoutes(r)

	log.Printf("🚀 Server running on :%s", port)
	http.ListenAndServe(":"+port, r)
}
