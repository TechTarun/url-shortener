package main

import (
	"log"
	"net/http"
	"url-shortener/internal/shortener"
	"url-shortener/internal/shortener/storage"
	"url-shortener/pkg/idgen"
)

func main() {
	repo := storage.NewInMemoryStore()
	generator := idgen.NewBase62Generator()
	service := shortener.NewService(repo, generator)
	handler := shortener.NewHandler(service)

	http.HandleFunc("/shorten", handler.Shorten)
	http.HandleFunc("/test", handler.Test)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
