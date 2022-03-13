package handlers

import (
	"net/http"
)

type Handler interface {
	GetRouter() http.Handler
	AllowCORS(allowedOrigins []string)
}
