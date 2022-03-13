package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func CreateCORSMiddleware(allowedOrigins []string) func(h http.Handler) http.Handler {
	cors := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
	})
	return cors.Handler
}
