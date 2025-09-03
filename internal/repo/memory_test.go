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
