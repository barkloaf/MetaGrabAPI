package main

import (
	"net/http"
	"sync"

	"github.com/barkloaf/MetaGrabAPI/misc"
	"golang.org/x/time/rate"
)

type ipRateLimiter struct {
	ips    map[string]*rate.Limiter
	mu     *sync.RWMutex
	rate   rate.Limit
	bucket int
}

func newIPRateLimiter(limit rate.Limit, bucket int) *ipRateLimiter {
	new := &ipRateLimiter{
		ips:    make(map[string]*rate.Limiter),
		mu:     &sync.RWMutex{},
		rate:   limit,
		bucket: bucket,
	}

	return new
}

func (lim *ipRateLimiter) addIP(ip string) *rate.Limiter {
	lim.mu.Lock()
	defer lim.mu.Unlock()

	limiter := rate.NewLimiter(lim.rate, lim.bucket)

	lim.ips[ip] = limiter

	return limiter
}

func (lim *ipRateLimiter) getLimiter(ip string) *rate.Limiter {
	lim.mu.Lock()
	limiter, exists := lim.ips[ip]

	if !exists {
		lim.mu.Unlock()
		return lim.addIP(ip)
	}

	lim.mu.Unlock()

	return limiter
}

var limiter = newIPRateLimiter(misc.Config.RateLimit, misc.Config.RateBucket)

func rateLimit(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		limiter := limiter.getLimiter(request.RemoteAddr)

		if !limiter.Allow() {
			http.Error(writer, http.StatusText(429), 429)
			return
		}

		handler.ServeHTTP(writer, request)
	})
}
