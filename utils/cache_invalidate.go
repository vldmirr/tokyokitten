// utils/cache_invalidate.go
package utils

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func InvalidateCache(redisClient *redis.Client, cacheKey string) {
	err := redisClient.Del(context.Background(), cacheKey).Err()
	if err != nil {
		log.Println(err)
	}
}