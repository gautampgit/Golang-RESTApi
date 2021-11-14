package httphandler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gautampgit/Golang-RESTApi/internal/comment"
	"github.com/gorilla/mux"
)

//Handler type stores pointer to the comments service
type Handler struct {
	Router  *mux.Router
	Service comment.Service
}

//SetupRoutes method sets all the routes of the app
func (h *Handler) SetupRoutes() {
	log.Println("Setting up th routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment/", h.GetAllComments).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comment/", h.PostComment).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods(http.MethodPut)
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods(http.MethodDelete)

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'm alive !")
	})
}

//NewHandler returns a pointer to a handler
func NewHandler(service comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

//GetComment - retrieves a comment by iD
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {

		log.Printf("Unable to parse ID to uint %v", err)
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprintf(w, "Unable to get all comments")
		log.Printf("Unable to get comments: %v", err)
	}

	fmt.Fprintf(w, "%+v", comment)
}

//GetAllComments- retrieves all comments from comments service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()

	if err != nil {
		fmt.Fprintf(w, "Unable to retrieve all comments")
		log.Printf("Unable to retrieve all comments %v", err)
	}
	fmt.Fprintf(w, "%+v", comments)
}

//PostComment - posts a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})
	if err != nil {
		fmt.Fprintf(w, "Unable to post a new comment")
	}
	fmt.Fprintf(w, "%+v", comment)
}

//UpdateComment - Updates an existing a comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/new",
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to Update the comment")
	}
	fmt.Fprintf(w, "%+v", comment)
}

//DeleteComment - deletes a comment based on ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse ID to uint")
	}
	err = h.Service.DeleteComment(uint(commentId))
	if err != nil {
		fmt.Fprintf(w, "Failed to delete comment")
	}
	fmt.Fprintf(w, "Successfully deleted comment")

}
