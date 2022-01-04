package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/bhanu-lab/Commentator/core/comment"
)

type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

type Response struct {
	Message string
}

// NewHandler -returns pointer to Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

/*
SetupRoutes - prepares handling paths and its functions
*/
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/health/check", h.HealthCheck)
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comments", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/{author}/{slug}/{body}", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/update/comment/{id}/{body}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/delete/comment/{id}", h.DeleteComment).Methods("DELETE")

}

/*
HealthCheck - deals with /health/check API
*/
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{"I am alive"}); err != nil {
		panic(err)
	}
}

/*
GetComment - gets comments for an ID given
*/
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := GetID(w, vars)
	if err != nil {
		return
	}
	comment, err := h.Service.GetComment(id)
	if err != nil {
		fmt.Fprintf(w, "error while getting comment")
	}
	fmt.Fprintf(w, "%+v", comment)
}

/*
GetComment - retrieves all comments
*/
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {

	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "error while getting all comments")
	}
	fmt.Fprintf(w, "%+v", comments)
}

/*
PostComment - inserts new comment into DB
*/
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	author := vars["author"]
	slug := vars["slug"]
	body := vars["body"]

	comment := comment.Comment{
		Slug:    slug,
		Body:    body,
		Author:  author,
		Created: time.Now(),
	}

	comment, err := h.Service.PostComment(comment)
	if err != nil {
		fmt.Fprint(w, "error while writing comment to DB")
	}

	fmt.Fprintf(w, "%+v", comment)

}

/*
UpdateComment - updates comment based on ID and Comment received
*/
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := GetID(w, vars)
	body := vars["body"]

	comm, err := h.Service.GetComment(id)
	if err != nil {
		fmt.Fprintf(w, "error while getting comm")
		return
	}
	comm.Body = body

	c, err := h.Service.UpdateComment(id, comm)
	if err != nil {
		fmt.Fprintf(w, "failed updating comm for if %d", id)
	}

	fmt.Fprintf(w, "%+v", c)
}

/*
DeleteComment - deletes cooments based on ID
*/
func (h *Handler) DeleteComment(w http.ResponseWriter, e *http.Request) {
	id, err := GetID(w, mux.Vars(e))
	if err != nil {
		return
	}
	h.Service.DeleteComment(id)
}

func GetID(w http.ResponseWriter, vars map[string]string) (int, error) {
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Fprintf(w, "error while converting id to int")
	}
	return int(id), err
}
