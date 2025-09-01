package handlers

import (
	"net/http"

	"github.com/diovch/microblog/internal/repo"
)

type PostHandler struct {
	r *repo.Repository
}

func NewPostHandler(r repo.Repository) *PostHandler {
	return &PostHandler{
		r: &r,
	}
}

func (p *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	
}

func (p *PostHandler) GetFeedHandler(w http.ResponseWriter, r *http.Request) {

}

func (p *PostHandler) LikePostHandler(w http.ResponseWriter, r *http.Request) {

}