package rest

import (
	"net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

// AntiAttackerHandler set anti attack server by limit request per ip
func (a *API) AntiAttackerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the IP address for the current user.
		ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			c.Status(http.StatusTooManyRequests)
			return
		}

		limiter := getUserByIP(ip)
		if !limiter.Allow() {
			c.Status(http.StatusTooManyRequests)
			return
		}
		c.Next()
	}
}

func getUserByIP(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 3)
		visitors[ip] = limiter
	}

	return limiter
}
