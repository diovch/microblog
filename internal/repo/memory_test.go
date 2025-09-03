package repo

import (
	"strconv"
	"testing"

	"github.com/diovch/microblog/internal/entity"
)

func TestCreateUser(t *testing.T) {
	repo := NewMemoryRepo()
	user := &entity.User{
		Username: "testuser",
	}
	err := repo.CreateUser(user.Username)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got, ok := repo.users[user.Username]
	if !ok {
		t.Fatalf("expected to find user, got error %v", err)
	}
	if got.Username != user.Username {
		t.Errorf("expected user %+v, got %+v", user, got)
	}
}

func BenchmarkCreateUser(b *testing.B) {
	repo := NewMemoryRepo()

	for i := 0; i < b.N; i++ {
		username := "user" + strconv.Itoa(i)
		err := repo.CreateUser(username)
		if err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
	}

	if len(repo.users) != b.N || repo.nextUserID != int64(b.N) {
		b.Errorf("expected %d users, got %d", b.N, len(repo.users))
	}
}

func TestCreatePost(t *testing.T) {
	repo := NewMemoryRepo()
	username := "testuser"
	err := repo.CreateUser(username)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	content := "Hello, world!"
	postID, err := repo.CreatePost(content, username)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if postID != 1 {
		t.Errorf("expected post ID 1, got %d", postID)
	}

	if len(repo.posts) != 1 {
		t.Fatalf("expected 1 post, got %d", len(repo.posts))
	}

	post := repo.posts[0]
	if post.ID != postID || post.Text != content || post.AuthorUsername != username {
		t.Errorf("post data mismatch: got %+v", post)
	}
}

func BenchmarkCreatePost(b *testing.B) {
	repo := NewMemoryRepo()
	username := "testuser"
	err := repo.CreateUser(username)
	if err != nil {
		b.Fatalf("expected no error, got %v", err)
	}

	for i := 0; i < b.N; i++ {
		content := "Post number " + strconv.Itoa(i)
		_, err := repo.CreatePost(content, username)
		if err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
	}

	if len(repo.posts) != b.N || repo.nextPostID != int64(b.N) {
		b.Errorf("expected %d posts, got %d", b.N, len(repo.posts))
	}
}

func TestLikePost(t *testing.T) {
	repo := NewMemoryRepo()
	username := "testuser"
	err := repo.CreateUser(username)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	content := "Hello, world!"
	postID, err := repo.CreatePost(content, username)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = repo.LikePost(postID, username)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	post := repo.posts[0]
	if len(post.Likes) != 1 || post.Likes[0] != username {
		t.Errorf("expected likes to contain %s, got %+v", username, post.Likes)
	}
}

func BenchmarkLikePost(b *testing.B) {
	repo := NewMemoryRepo()
	username := "testuser"
	err := repo.CreateUser(username)
	if err != nil {
		b.Fatalf("expected no error, got %v", err)
	}

	content := "Hello, world!"
	postID, err := repo.CreatePost(content, username)
	if err != nil {
		b.Fatalf("expected no error, got %v", err)
	}

	for i := 0; i < b.N; i++ {
		newUserName := username + strconv.Itoa(i)
		err := repo.CreateUser(newUserName)
		if err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
		err = repo.LikePost(postID, username+strconv.Itoa(i))
		if err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
	}

	post := repo.posts[0]
	if len(post.Likes) != b.N {
		b.Errorf("expected %d likes, got %d", b.N, len(post.Likes))
	}
}
