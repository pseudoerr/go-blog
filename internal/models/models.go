package models

import (
	"time"
)

type Post struct {
	PostID    int64     `json:"post_id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Tags      []Tag     `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	UserID         int64     `json:"user_id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	TimeStamp      time.Time `json:"time_stamp"`
}

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type PostTag struct {
	PostID int64 `json:"post_id"`
	TagID  int64 `json:"tag_id"`
}
