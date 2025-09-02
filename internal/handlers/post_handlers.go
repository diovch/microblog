package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/diovch/microblog/internal/entity"
	"github.com/diovch/microblog/internal/repo"
)

type PostHandler struct {
	r repo.Repository
	validator
}

func NewPostHandler(r repo.Repository) *PostHandler {
	return &PostHandler{
		r: r,
	}
}

func (p *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if err := p.ValidateJsonContentType(r); err != nil {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	var post entity.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	createdPost, err := p.r.CreatePost(post.Text, post.AuthorUsername)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)
}

func (p *PostHandler) GetFeedHandler(w http.ResponseWriter, r *http.Request) {
	posts := p.r.GetAllPosts()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (p *PostHandler) LikePostHandler(w http.ResponseWriter, r *http.Request) {

}
