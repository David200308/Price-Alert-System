package initializers

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:86379",
		Password: "",
		DB:       0,
	})
	result := rdb.Ping(context.Background())
	fmt.Println("redis ping:", result.Val())
	if result.Val() != "PONG" {
		fmt.Println("Failed to connect to redis")
		return
	}
}
