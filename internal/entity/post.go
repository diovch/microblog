package entity

type Post struct {
	ID             int64    `json:"id"`
	Text           string   `json:"text"`
	AuthorUsername string   `json:"author_username"`
	Likes          []string `json:"likes"`
}
