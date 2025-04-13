package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log sebelum request diproses
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
		
		// Proses request
		next.ServeHTTP(w, r)
		
		// Log setelah request selesai diproses
		log.Printf("Request completed: %s %s", r.Method, r.URL.Path)
	})
}
