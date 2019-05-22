package utilities

import (
	"net/http"

	"golang.org/x/time/rate"
)

// Limiter ...
var Limiter = rate.NewLimiter(1, 4)

// Limit ...
func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if Limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
