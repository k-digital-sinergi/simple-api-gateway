package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v9"
	"log"
	"net/http"
	"simple-api-gateway/connection"
	"simple-api-gateway/pkg/util"
)

const (
	rateLimit = 3
)

func RateLimitMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		limiter := redis_rate.NewLimiter(connection.GetRedisConnection(ctx))
		result, err := limiter.Allow(ctx, ip, redis_rate.PerMinute(rateLimit))
		if err != nil {
			log.Panic(err)
		}

		if result.Allowed == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, util.ErrorResponse(errors.New("rate limit is exhausted")))
			return
		}

		ctx.Next()
	}
}
