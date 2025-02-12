package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiter manages rate limiting per visitor
type RateLimiter struct {
	sync.RWMutex
	visitors map[string]*rate.Limiter
	rate     rate.Limit
	burst    int
	ttl      time.Duration
	lastSeen map[string]time.Time
}

// NewRateLimiter creates a new rate limiter with specified requests/second and burst
func NewRateLimiter(requestsPerSecond float64, burst int, ttl time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*rate.Limiter),
		lastSeen: make(map[string]time.Time),
		rate:     rate.Limit(requestsPerSecond),
		burst:    burst,
		ttl:      ttl,
	}

	go rl.cleanupVisitors()
	return rl
}

// getVisitor retrieves or creates a rate limiter for a visitor
func (rl *RateLimiter) getVisitor(ip string) *rate.Limiter {
	rl.Lock()
	defer rl.Unlock()

	limiter, exists := rl.visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(rl.rate, rl.burst)
		rl.visitors[ip] = limiter
	}

	rl.lastSeen[ip] = time.Now()
	return limiter
}

// cleanupVisitors removes rate limiters for IPs that haven't been seen recently
func (rl *RateLimiter) cleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		rl.Lock()
		for ip, lastSeen := range rl.lastSeen {
			if time.Since(lastSeen) > rl.ttl {
				delete(rl.visitors, ip)
				delete(rl.lastSeen, ip)
			}
		}
		rl.Unlock()
	}
}

// RateLimit middleware for Gin
func (rl *RateLimiter) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		limiter := rl.getVisitor(ip)

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":       "rate limit exceeded",
				"retry_after": time.Duration(1/float64(rl.rate)) * time.Second,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
