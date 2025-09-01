package handlers

import (
	"net/http"

	"github.com/diovch/microblog/internal/repo"
)

type UserHandler struct {
	r *repo.Repository
}

func NewUserHandler(r repo.Repository) *UserHandler {
	return &UserHandler{
		r: &r,
	}
}

func (u *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {

}

