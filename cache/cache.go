package cache

import (
	"github.com/go-redis/redis/v7"
	"log"
	"os"
	"strconv"
)

var Client *redis.Client

func init() {
	dbIdx, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("failed to connect to redis server: REDIS_DB environment variable is not available: %v", err)
	}

	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbIdx,
	})

	_, err = Client.Ping().Result()
	if err != nil {
		log.Fatalf("failed to connect to redis server: %v", err)
	}
}
