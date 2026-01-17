package storage

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
	ctx    context.Context
	ttl    time.Duration
}

func NewRedisStore(ttl time.Duration) *RedisStore {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	rdb := redis.NewClient(&redis.Options{Addr: addr})
	return &RedisStore{
		client: rdb,
		ctx:    context.Background(),
		ttl:    ttl,
	}
}

func (r *RedisStore) Save(shortCode string, longUrl string) error {
	return r.client.Set(r.ctx, shortCode, longUrl, r.ttl).Err()
}

func (r *RedisStore) Get(shortCode string) (string, error) {
	value, error := r.client.Get(r.ctx, shortCode).Result()
	if error != nil {
		return "", error
	}
	return value, nil
}
