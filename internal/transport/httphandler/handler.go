package httphandler

import (
	"encoding/json"
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

//Response - struct to store the response
type Response struct {
	Message string
	Error   string
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
		convertJSON(w, http.StatusOK, Response{Message: "I'm alive! "})
	})
}

//NewHandler returns a pointer to a handler
func NewHandler(service comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

//convertJSON converts the response object to JSON
func convertJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}

//GetComment - retrieves a comment by iD
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse ID to uint ", err)
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Unable to get all comments", err)
		return
	}
	convertJSON(w, http.StatusOK, comment)
}

//GetAllComments- retrieves all comments from comments service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()

	if err != nil {
		sendErrorResponse(w, "Unable to retrieve all comments", err)
		return
	}
	convertJSON(w, http.StatusOK, comments)
}

//PostComment - posts a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		sendErrorResponse(w, "Failed to convert body to JSON", err)
		return
	}

	comment, err = h.Service.PostComment(comment)
	if err != nil {
		sendErrorResponse(w, "Unable to post a new comment", err)
		return
	}
	convertJSON(w, http.StatusOK, comment)
}

//UpdateComment - Updates an existing a comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse ID to uint", err)
		return
	}
	var comment comment.Comment

	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		sendErrorResponse(w, "Failed to convert body to JSON", err)
		return
	}

	comment, err = h.Service.UpdateComment(uint(commentId), comment)
	if err != nil {
		sendErrorResponse(w, "Failed to Update the comment", err)
		return
	}
	convertJSON(w, http.StatusOK, comment)
}

//DeleteComment - deletes a comment based on ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse ID to uint", err)
		return
	}
	err = h.Service.DeleteComment(uint(commentId))
	if err != nil {
		sendErrorResponse(w, "Failed to delete comment", err)
		return
	}
	convertJSON(w, http.StatusOK, Response{Message: "Successfully Deleted !..."})

}
