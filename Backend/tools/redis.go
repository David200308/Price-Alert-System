package tools

import (
	"context"
	"time"

	"github.com/David200308/go-api/Backend/initializers"
	"github.com/go-redis/redis/v8"
)

func CacheSet(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	statusCmd := initializers.Redis.Set(ctx, key, value, expiration)

	if err := statusCmd.Err(); err != nil {
		return err
	}
	return nil
}

func CacheGet(ctx context.Context, key string) (string, error) {
	resultCmd := initializers.Redis.Get(ctx, key)

	if err := resultCmd.Err(); err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}

	return resultCmd.Val(), nil
}
