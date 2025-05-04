package main

import (
	"github.com/maksiex/dir/internal/configs"
	"github.com/maksiex/dir/pkg/db"
	"log"
	"net/http"
)

func main() {
	configs.LoadInitialConfig()
	dbConfig := configs.LoadDBConfig()
	db.DirDb(dbConfig)

	log.Println("✅ Starting backend...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Backend is running 🚀"))
		if err != nil {
			log.Println(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
