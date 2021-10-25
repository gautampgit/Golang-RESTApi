package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//HAndler type stores pointer to the comments service
type Handler struct {
	Router *mux.Router
}

//NewHandler returns a pointer to a handler
func NewHandler() *Handler {
	return &Handler{}
}

//SetupRoutes method sets all the routes of the app
func (h *Handler) SetupRoutes() {
	log.Println("Setting up th routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'm alive !")
	})
}
