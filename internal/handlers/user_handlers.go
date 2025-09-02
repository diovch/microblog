package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/diovch/microblog/internal/entity"
	"github.com/diovch/microblog/internal/repo"
)

type UserHandler struct {
	r repo.Repository
}

func NewUserHandler(r repo.Repository) *UserHandler {
	return &UserHandler{
		r: r,
	}
}

func (u *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	content_type := r.Header.Get("Content-Type")

	if content_type != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := u.r.CreateUser(user.Username); err != nil {
		http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
