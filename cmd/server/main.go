package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"url-shortener/internal/shortener"
	"url-shortener/internal/shortener/storage"
	"url-shortener/pkg/idgen"

	"github.com/joho/godotenv"
)

func main() {
	// repo := storage.NewInMemoryStore()
	_ = godotenv.Load()
	ttlStr := os.Getenv("REDIS_TTL")
	ttl, err := time.ParseDuration(ttlStr)
	if err != nil {
		ttl = 24 * time.Hour
	}
	shortCodeLength, _ := strconv.Atoi(os.Getenv("SHORT_CODE_LENGTH"))
	nodeId, _ := strconv.ParseInt(os.Getenv("NODE_ID"), 10, 64)

	repo := storage.NewRedisStore(ttl)
	generator, _ := idgen.NewSnowflakeGenerator(shortCodeLength, nodeId)
	service := shortener.NewService(repo, generator)
	handler := shortener.NewHandler(service)

	http.HandleFunc("/shorten", handler.Shorten)
	http.HandleFunc("/test", handler.Test)
	http.HandleFunc("/", handler.Resolve)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
