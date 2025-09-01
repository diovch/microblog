package repo

import "github.com/diovch/microblog/internal/entity"

type Repository interface {
	// User methods
	CreateUser(username, password string) error
	GetUserByUsername(username string) (*entity.User, error)

	// Post methods
	CreatePost(content string, authorID int64) (int64, error)
	GetAllPosts() ([]*entity.Post, error)
	LikePost(id int64) error
}
