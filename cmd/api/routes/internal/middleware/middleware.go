package middleware

import (
	"log"
	"net/http"
	"runtime"
)

// recovers from unexpected panic.
func RecoveryFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 2048)
				_ = runtime.Stack(buf, false)
				log.Printf("[PANIC]: %s\n", err)
				log.Printf("[STACK TRACE]: %s\n", buf)
				log.Println("[INFO]:Recoverd.")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// sets 'application/json' as default content type
func SetContentTypeJsonFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
