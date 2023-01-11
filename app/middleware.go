package app

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, request *http.Request) {
			flag := true
			fmt.Println("Checking authentication")
			if !flag {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, http.StatusText(http.StatusUnauthorized))
				return
			}

			next(w, request)
		}
	}
}

func Logging() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, request *http.Request) {
			start := time.Now()
			defer log.Println(request.URL.Path, time.Since(start))
			next(w, request)
		}
	}
}
