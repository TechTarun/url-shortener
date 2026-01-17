package main

import (
	"log"
	"net/http"
	"url-shortener/internal/shortener"
	"url-shortener/internal/shortener/storage"
	"url-shortener/pkg/idgen"
)

func main() {
	// repo := storage.NewInMemoryStore()
	repo := storage.NewRedisStore("localhost:6379", 0)
	generator, _ := idgen.NewSnowflakeGenerator(8, 1)
	service := shortener.NewService(repo, generator)
	handler := shortener.NewHandler(service)

	http.HandleFunc("/shorten", handler.Shorten)
	http.HandleFunc("/test", handler.Test)
	http.HandleFunc("/", handler.Resolve)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
