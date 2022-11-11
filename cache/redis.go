package cache

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var Client	*redis.Client


func Init() {
	redisURL := os.Getenv("REDIS_SERVER_URL")

	options := &redis.Options{
		Addr:     redisURL,
		Password: "",
	}
	Client = redis.NewClient(options)
}

func Add(key string, value string) {
	ctx := context.TODO()

	Client.Append(ctx, key, value).Result()
}

func Get(key string) []string {
	ctx := context.TODO()

	messages, _ := Client.LRange(ctx, key, -1000, 1000).Result()
	return messages
}
