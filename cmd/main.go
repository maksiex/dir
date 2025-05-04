package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/maksiex/dir/internal/configs"
	"github.com/maksiex/dir/pkg/db"
	"io"
	"log"
	"net/http"
)

func main() {
	configs.LoadInitialConfig()
	dbConfig := configs.LoadDBConfig()
	db.DirDb(dbConfig)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
	}))
	r.Use(middleware.RequestID)

	r.Get("/airports", func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest("GET", "https://api.aviationapi.com/v1/airports?apt=ART&group=1", nil)
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Failed to fetch data", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			log.Println("Failed to write response:", err)
		}
	})

	log.Println("✅ API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
