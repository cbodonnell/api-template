package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cbodonnell/api-template/config"
	"github.com/cbodonnell/api-template/services"
	"github.com/gorilla/mux"
)

type MuxHandler struct {
	conf    config.Configuration
	service services.Service
	router  *mux.Router
}

var (
	nameParam = "name"
)

func NewMuxHandler(_config config.Configuration, _service services.Service) Handler {
	h := &MuxHandler{
		conf:    _config,
		service: _service,
		router:  mux.NewRouter(),
	}
	h.setupRoutes()
	return h
}

func (h *MuxHandler) setupRoutes() {

	get := h.router.NewRoute().Subrouter() // -> public GET requests
	// get.Use(middleware)
	get.HandleFunc(fmt.Sprintf("/users/{%s:[[:alnum:]]+}", nameParam), h.GetUser).Methods("GET", "OPTIONS")

	sink := h.router.NewRoute().Subrouter() // -> sink to handle all other routes
	sink.PathPrefix("/").HandlerFunc(h.SinkHandler).Methods("GET", "OPTIONS")

}

func (h *MuxHandler) GetRouter() http.Handler {
	return h.router
}

func (h *MuxHandler) AllowCORS(allowedOrigins []string) {
	// cors := CreateCORSMiddleware(allowedOrigins)
	// h.router.Use(cors)
}

func (h *MuxHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)[nameParam]
	user, err := h.service.GetUserByName(name)
	if err != nil {
		// h.response.NotFound(w, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// actor := h.resource.GenerateActor(user.Name)
	// w.Header().Set("Content-Type", activitypub.ContentType)
	json.NewEncoder(w).Encode(user)
}

func (h *MuxHandler) SinkHandler(w http.ResponseWriter, r *http.Request) {
	// h.response.NotFound(w, fmt.Errorf("endpoint %s does not exist", r.URL))
}
