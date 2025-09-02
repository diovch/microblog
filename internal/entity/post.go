package entity

type Post struct {
	ID             int64    `json:"id,omitempty"`
	Text           string   `json:"text"`
	AuthorUsername string   `json:"author_username"`
	Likes          []string `json:"likes"`
}

func NewPost(text string, authorUsername string) *Post {
	return &Post{
		Text:           text,
		AuthorUsername: authorUsername,
		Likes:          make([]string, 0),
	}
}

func (p *Post) Like(username string) {
	p.Likes = append(p.Likes, username)
}