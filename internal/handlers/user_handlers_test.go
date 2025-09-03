package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/diovch/microblog/internal/entity"
	"github.com/diovch/microblog/internal/logger"
)

// --- Mocks ---

type mockRepo struct {
	createUserFunc        func(username string) error
	getUserByUsernameFunc func(username string) (*entity.User, error)
	createPostFunc        func(content, authorUsername string) (int64, error)
	getAllPostsFunc       func() []*entity.Post
	likePostFunc          func(id int64, username string) error
}

func (m *mockRepo) CreateUser(username string) error {
	return m.createUserFunc(username)
}

func (m *mockRepo) GetUserByUsername(username string) (*entity.User, error) {
	return m.getUserByUsernameFunc(username)
}

func (m *mockRepo) CreatePost(content, authorUsername string) (int64, error) {
	return m.createPostFunc(content, authorUsername)
}
func (m *mockRepo) GetAllPosts() []*entity.Post {
	return m.getAllPostsFunc()
}
func (m *mockRepo) LikePost(id int64, username string) error {
	return m.likePostFunc(id, username)
}

// --- Test Helper ---

func NewLogger() *logger.Logger {
	return logger.NewLogger(1)
}

// --- Tests ---

func TestRegisterHandler_Success(t *testing.T) {
	handler := &UserHandler{
		l: NewLogger(),
		r: &mockRepo{
			createUserFunc: func(username string) error { return nil },
		},
	}

	jbody, _ := json.Marshal(map[string]string{"Username": "testuser"})
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(jbody))
	
	defer req.Body.Close()
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.RegisterHandler(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("status is wrong %d", rr.Code)
	}
}
