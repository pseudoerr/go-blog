package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Post struct {
	PostID    int64     `json:"post_id" db:"post_id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	Body      string    `json:"body" db:"body"`
	Tags      []Tag     `json:"tags" db:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type User struct {
	UserID         uuid.UUID `json:"user_id" db:"user_id"`
	Username       string    `json:"username" db:"username"`
	HashedPassword string    `json:"-" db:"hashed_password"`
	TimeStamp      time.Time `json:"time_stamp" db:"time_stamp"`
}

type Tag struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type PostTag struct {
	PostID int64 `json:"post_id" db:"post_id"`
	TagID  int64 `json:"tag_id" db:"tag_id"`
}
