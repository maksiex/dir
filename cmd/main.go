package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/maksiex/dir/internal/configs"
	dbpkg "github.com/maksiex/dir/pkg/db"
	"github.com/maksiex/dir/pkg/models"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	configs.LoadInitialConfig()
	dbConfig := configs.LoadDBConfig()
	db := dbpkg.DirDb(dbConfig).Gorm

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
	}))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	apiKeys := map[string]string{
		"gnews":      os.Getenv("GNEWS_API_KEY"),
		"mediastack": os.Getenv("MEDIASTACK_API_KEY"),
		"currents":   os.Getenv("CURRENTS_API_KEY"),
		"newsapi":    os.Getenv("NEWS_API_KEY"),
	}

	r.Get("/fetch/gnews", fetchGnewsHandler(db, apiKeys["gnews"]))
	r.Get("/fetch/mediastack", fetchMediastackHandler(db, apiKeys["mediastack"]))
	r.Get("/fetch/currents", fetchCurrentsHandler(db, apiKeys["currents"]))
	r.Get("/fetch/newsapi", fetchNewsapiHandler(db, apiKeys["newsapi"]))

	log.Println("✅ API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func fetchGnewsHandler(db *gorm.DB, apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(fmt.Sprintf("https://gnews.io/api/v4/top-headlines?category=world&lang=en&country=us&max=25&apikey=%s", apiKey))
		if err != nil {
			http.Error(w, "Failed to fetch GNews data", http.StatusInternalServerError)
			log.Println("GNews fetch error:", err)
			return
		}
		defer resp.Body.Close()

		var result struct {
			Articles []models.NewsFromGnews `json:"articles"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			http.Error(w, "Failed to decode GNews response", http.StatusInternalServerError)
			log.Println("GNews decode error:", err)
			return
		}

		for _, article := range result.Articles {
			if err := db.Create(&article).Error; err != nil {
				log.Println("DB insert error:", err)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("✅ GNews articles saved successfully"))
	}
}

func fetchMediastackHandler(db *gorm.DB, apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("http://api.mediastack.com/v1/news?access_key=%s&languages=en", apiKey)

		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Failed to fetch Mediastack data", http.StatusInternalServerError)
			log.Println("Mediastack fetch error:", err)
			return
		}
		defer resp.Body.Close()

		var result struct {
			Data []models.NewsFromMediastack `json:"data"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			http.Error(w, "Failed to decode Mediastack response", http.StatusInternalServerError)
			log.Println("Mediastack decode error:", err)
			return
		}

		for _, article := range result.Data {
			if err := db.Create(&article).Error; err != nil {
				log.Println("DB insert error (mediastack):", err)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("✅ Mediastack articles saved"))
	}
}

func fetchCurrentsHandler(db *gorm.DB, apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("https://api.currentsapi.services/v1/latest-news?apiKey=%s", apiKey)

		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Failed to fetch Currents data", http.StatusInternalServerError)
			log.Println("Currents fetch error:", err)
			return
		}
		defer resp.Body.Close()

		var result struct {
			News []models.NewsFromCurrents `json:"news"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			http.Error(w, "Failed to decode Currents response", http.StatusInternalServerError)
			log.Println("Currents decode error:", err)
			return
		}

		for _, article := range result.News {
			if err := db.Create(&article).Error; err != nil {
				log.Println("DB insert error (currents):", err)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("✅ Currents articles saved"))
	}
}

func fetchNewsapiHandler(db *gorm.DB, apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&category=business&apiKey=%s", apiKey)

		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Failed to fetch NewsAPI data", http.StatusInternalServerError)
			log.Println("NewsAPI fetch error:", err)
			return
		}
		defer resp.Body.Close()

		var result struct {
			Articles []models.NewsFromNewsapi `json:"articles"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			http.Error(w, "Failed to decode NewsAPI response", http.StatusInternalServerError)
			log.Println("NewsAPI decode error:", err)
			return
		}

		for _, article := range result.Articles {
			if err := db.Create(&article).Error; err != nil {
				log.Println("DB insert error (newsapi):", err)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("✅ NewsAPI articles saved"))
	}
}
