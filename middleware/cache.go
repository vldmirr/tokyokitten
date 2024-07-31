// middleware/cache.go
package middleware

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Cache(ttl time.Duration, redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		cacheKey := c.Request.Method + ":" + c.Request.URL.Path
		val, err := redisClient.Get(context.Background(), cacheKey).Result()
		if err != nil {
			log.Println(err)
			c.Next()
			return
		}

		var data interface{}
		err = json.Unmarshal([]byte(val), &data)
		if err != nil {
			log.Println(err)
			c.Next()
			return
		}

		c.JSON(200, data)
		c.Abort()
	}
}