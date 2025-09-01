package repo

import (
	"testing"

	"github.com/diovch/microblog/internal/entity"
)

func TestCreateUser(t *testing.T) {
	repo := NewMemoryRepo()
	user := &entity.User{Username: "alice"}

	err := repo.CreateUser(user)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Попытка создать пользователя с тем же именем
	err = repo.CreateUser(&entity.User{Username: "alice"})
	if err == nil {
		t.Fatal("expected error for duplicate username, got nil")
	}
}

func TestGetUserByUsername(t *testing.T) {
	repo := NewMemoryRepo()
	user := &entity.User{Username: "bob"}
	_ = repo.CreateUser(user)

	got, err := repo.GetUserByUsername("bob")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Username != "bob" {
		t.Errorf("expected username 'bob', got '%s'", got.Username)
	}

	_, err = repo.GetUserByUsername("notfound")
	if err == nil {
		t.Fatal("expected error for not found user, got nil")
	}
}

func TestCreatePost(t *testing.T) {
	repo := NewMemoryRepo()
	post := &entity.Post{Text: "hello", AuthorUsername: "alice"}

	id, err := repo.CreatePost(post)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != 1 {
		t.Errorf("expected post ID 1, got %d", id)
	}
}

func TestGetAllPosts(t *testing.T) {
	repo := NewMemoryRepo()
	repo.CreatePost(&entity.Post{Text: "post1", AuthorUsername: "alice"})
	repo.CreatePost(&entity.Post{Text: "post2", AuthorUsername: "bob"})

	posts := repo.GetAllPosts(0)
	if len(posts) != 2 {
		t.Errorf("expected 2 posts, got %d", len(posts))
	}
}

func TestLikePost(t *testing.T) {
	repo := NewMemoryRepo()
	post := &entity.Post{Text: "like me", AuthorUsername: "alice"}
	id, _ := repo.CreatePost(post)

	err := repo.LikePost(id, "bob")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	posts := repo.GetAllPosts(0)
	if len(posts[0].Likes) != 1 || posts[0].Likes[0] != "bob" {
		t.Errorf("expected like from 'bob', got %+v", posts[0].Likes)
	}

	// Лайк несуществующего поста
	err = repo.LikePost(999, "bob")
	if err == nil {
		t.Fatal("expected error for non-existent post, got nil")
	}
}
