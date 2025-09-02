package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/diovch/microblog/internal/entity"
	"github.com/diovch/microblog/internal/logger"
	"github.com/diovch/microblog/internal/repo"
)

type UserHandler struct {
	r repo.Repository
	l *logger.Logger
	validator
}

func NewUserHandler(r repo.Repository, l *logger.Logger) *UserHandler {
	return &UserHandler{
		r: r,
		l: l,
	}
}

func (u *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if err := u.ValidateJsonContentType(r); err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
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
	u.l.LogInfo("User " + user.Username + " registered successfully")
}
