package repo

import (
	"errors"
	"sync"

	"github.com/diovch/microblog/internal/entity"
)

type MemoryRepo struct {
	users map[string]*entity.User
	posts []*entity.Post

	nextUserID int64
	nextPostID int64

	userMutex *sync.Mutex
	postMutex *sync.Mutex
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		users:      make(map[string]*entity.User),
		posts:      make([]*entity.Post, 0),
		nextUserID: 0,
		nextPostID: 0,
		userMutex:  &sync.Mutex{},
		postMutex:  &sync.Mutex{},
	}
}

func (r *MemoryRepo) CreateUser(user *entity.User) error {
	r.userMutex.Lock()
	defer r.userMutex.Unlock()

	if _, exists := r.users[user.Username]; exists {
		return errors.New("username already exists")
	}

	r.nextUserID++
	user.ID = r.nextUserID
	r.users[user.Username] = user

	return nil
}

func (r *MemoryRepo) GetUserByUsername(username string) (*entity.User, error) {
	r.userMutex.Lock()
	defer r.userMutex.Unlock()

	user, exists := r.users[username]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *MemoryRepo) CreatePost(post *entity.Post) (int64, error) {
	r.postMutex.Lock()
	defer r.postMutex.Unlock()

	r.nextPostID++
	post.ID = r.nextPostID
	r.posts = append(r.posts, post)

	return post.ID, nil
}

func (r *MemoryRepo) GetAllPosts(id int64) []*entity.Post {
	r.postMutex.Lock()
	defer r.postMutex.Unlock()

	res := make([]*entity.Post, 0, len(r.posts))
	copy(res, r.posts)

	return res
}

func (r *MemoryRepo) LikePost(id int64, username string) error {
	r.postMutex.Lock()
	defer r.postMutex.Unlock()

	for _, post := range r.posts {
		if post.ID == id {
			post.Likes = append(post.Likes, username)
			return nil
		}
	}

	return errors.New("post not found")
}
