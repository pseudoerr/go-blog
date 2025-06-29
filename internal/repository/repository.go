package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pseudoerr/go-blog/internal/models"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post models.Post) (models.Post, error)
	GetPostById(ctx context.Context, postID int64) (models.Post, error)
	ListPosts(ctx context.Context, userID int64, limit, offset int) ([]models.Post, error)
	UpdatePost(ctx context.Context, post models.Post) (models.Post, error)
	DeletePost(ctx context.Context, postID int64) error
	ListPostsByTag(ctx context.Context, tag string) ([]models.Post, error)
}

type TagRepository interface {
	GetOrCreateTag(ctx context.Context, tag models.Tag) (models.Tag, error)
	ListTags(ctx context.Context, postID int64) ([]models.Tag, error)
	ListAllTags(ctx context.Context) ([]models.Tag, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, username string, hashedPassword string) (models.User, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	GetUserByID(ctx context.Context, userID int64) (models.User, error)
}

type Repository struct {
	PostRepository
	TagRepository
	UserRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PostRepository:
		TagRepository: 
		UserRepository:
	}
}
