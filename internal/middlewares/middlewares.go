package middlewares

import (
	"github.com/gabrielix29/go-rest-api/pkg/logger"
	"net/http"
)

// JSON middleware to set content-type header for json
func Json(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug(r.Method, " ", r.URL)
		next.ServeHTTP(w, r)
	})
}
