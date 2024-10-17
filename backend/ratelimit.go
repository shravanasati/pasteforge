package main

import (
	"time"

	"github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
)

// NewRateLimiter returns a gin middleware for rate limiting. 
// rps is the requests per second.
func NewRateLimiter(limit uint, rate time.Duration) gin.HandlerFunc {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  rate,
		Limit: limit,
	})
	middleware := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: func(ctx *gin.Context, info ratelimit.Info) {
			ctx.JSON(429, gin.H{
				"error": "Too many requests, try again in " + time.Until(info.ResetTime).String(),
			})
			ctx.Abort()
		},
		KeyFunc: func(ctx *gin.Context) string {
			return ctx.ClientIP()
		},
	})

	return middleware
}
