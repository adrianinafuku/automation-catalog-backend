package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"automation-catalog-backend/internal/n8n"
	"automation-catalog-backend/internal/supabase"
)

func RegisterRoutes(r chi.Router) {
	r.Get("/health", healthHandler)
	r.Get("/automations", automationsHandler)
	r.Post("/demo/{id}", demoHandler)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func automationsHandler(w http.ResponseWriter, r *http.Request) {
	list, err := supabase.ListAutomations()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(list)
}

func demoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := n8n.TriggerDemo(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("demo executed"))
}
