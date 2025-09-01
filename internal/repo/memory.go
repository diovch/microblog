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

	userMutex *sync.RWMutex
	postMutex *sync.RWMutex
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		users:      make(map[string]*entity.User),
		posts:      make([]*entity.Post, 0),
		nextUserID: 0,
		nextPostID: 0,
		userMutex:  &sync.RWMutex{},
		postMutex:  &sync.RWMutex{},
	}
}

func (r *MemoryRepo) CreateUser(username string) error {
	r.userMutex.Lock()
	defer r.userMutex.Unlock()

	if _, exists := r.users[username]; exists {
		return errors.New("username already exists")
	}

	user := entity.NewUser(username)
	r.nextUserID++
	user.ID = r.nextUserID
	r.users[user.Username] = user

	return nil
}

func (r *MemoryRepo) GetUserByUsername(username string) (*entity.User, error) {
	r.userMutex.RLock()
	defer r.userMutex.RUnlock()

	user, exists := r.users[username]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *MemoryRepo) CreatePost(content string, authorUsername string) (int64, error) {
	r.postMutex.Lock()
	defer r.postMutex.Unlock()

	post := entity.NewPost(content, authorUsername)
	r.nextPostID++
	post.ID = r.nextPostID
	r.posts = append(r.posts, post)

	return post.ID, nil
}


func (r *MemoryRepo) GetAllPosts() []*entity.Post {
	r.postMutex.RLock()
	defer r.postMutex.RUnlock()

	return r.posts[:]
}


func (r *MemoryRepo) LikePost(id int64, username string) error {
	r.postMutex.RLock()
	defer r.postMutex.RUnlock()

	if id <= 0 || id > r.nextPostID {
		return errors.New("id is not valid")
	}

	if r.posts[id-1].ID != id {
		return errors.New("post not found")
	}

	post := r.posts[id-1]
	post.Like(username)
	return nil
}
