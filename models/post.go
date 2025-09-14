package models

import (
	"time"
)

// 投稿情報
type Post struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostWithComments struct {
	Post     Post      `json:"post"`
	Comments []Comment `json:"comments"`
}
