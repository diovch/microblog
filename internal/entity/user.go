package entity

type User struct {
	ID       int64  `json:"id,omitempty"`
	Username string `json:"username"`
}

func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}