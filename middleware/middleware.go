package middleware

import (
	"context"
	"fmt"
	"time"

	"net/http"

	apicalls "github.com/lancer2672/DandelionServer_Go/api_calls"
	"github.com/lancer2672/DandelionServer_Go/utils"
	"github.com/rs/zerolog/log"
)

func CheckApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//check permission to add movie

		value := context.Background().Value("payload")
		fmt.Println("CONTEXT VLUE", value)
		apiKey := r.Header.Get("x-api-key")
		fmt.Println("Http Api Header:", apiKey)
		if r.Method == "POST" && r.URL.Path == "/movies" {
			role, permission, err := apicalls.ExternalServiceIns.CheckApiKey(apiKey)
			if err != nil {
				http.Error(w, "Invalid API Key", http.StatusUnauthorized)
				return
			}
			// editor_api_key
			// admin_api_key
			// user_api_key
			if apiKey == "admin_api_key" && utils.StringInSlice("movie", permission.Write) {
				fmt.Println("Permission granted")
			} else {
				http.Error(w, "Permission denied", http.StatusUnauthorized)
				return
			}
			fmt.Println("ROLE", role)
			fmt.Println("PERMISSION", permission)
		}

		next.ServeHTTP(w, r)
	})
}

type CustomResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		resWriter := &CustomResponseWriter{w, http.StatusOK}
		handler.ServeHTTP(resWriter, r)
		log.Info().Str("METHOD", r.Method).Str("PATH", r.URL.Path).Dur("duration", time.Since(start)).Int("CODE", resWriter.statusCode).Str("STATUS", http.StatusText(resWriter.statusCode)).Msg("HTTP Request")

	})
}
