package cache

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Cli	*redis.Client
}

var Client Cache

func Init() {
	redisURL := os.Getenv("REDIS_SERVER_URL")

	options := &redis.Options{
		Addr:     redisURL,
		Password: "",
	}
	Client.Cli = redis.NewClient(options)
}

func (c *Cache) Add(key string, value interface{}) {
	ctx := context.TODO()

	c.Cli.LPush(ctx, key, value).Result()
}

func (c *Cache) Get(key string) []string {
	ctx := context.TODO()

	messages, _ := c.Cli.LRange(ctx, key, 0, -1).Result()
	return messages
}
