package server

import (
	"fmt"
	"net/http"
	"time"
)

func AccessLoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("access login middleware", r.URL.Path)
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("[%s] %s, %s %s\n", r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}
