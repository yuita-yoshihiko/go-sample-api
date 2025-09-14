package models

import (
	"time"
)

// ユーザー情報
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserWithPosts struct {
	User  User   `json:"user"`
	Posts []Post `json:"posts"`
}
