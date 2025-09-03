package repo

import "github.com/diovch/microblog/internal/entity"

type Repository interface {
	// User methods
	CreateUser(username string) error
	GetUserByUsername(username string) (*entity.User, error)

	// Post methods
	CreatePost(content, authorUsername string) (int64, error)
	GetAllPosts() ([]*entity.Post)
	LikePost(id int64, username string) error
}
