package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pseudoerr/go-blog/internal/models"
)

type tagRepository struct {
	db *sqlx.DB
}

func NewTagRepository(db *sqlx.DB) TagRepository {
	return &tagRepository{
		db: db,
	}
}
func (r *tagRepository) GetOrCreateTag(ctx context.Context, tag models.Tag) (models.Tag, error) {
	query := `
	INSERT INTO tags (name)
	VALUES ($1)
	ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name
	RETURNING id, name
	`
	var t models.Tag
	err := r.db.QueryRowxContext(ctx, query, tag.Name).StructScan(&t)
	if err != nil {
		return models.Tag{}, fmt.Errorf("GetOrCreateTag: %w", err)
	}
	return t, nil
}

func (r *tagRepository) ListTags(ctx context.Context, postID int64) ([]models.Tag, error) {
	query := `
	SELECT t.id, t.name
	FROM tags t
	JOIN post_tags pt ON t.id = pt.tag_id
	WHERE pt.post_id = $1
	`
	var tags []models.Tag
	err := r.db.SelectContext(ctx, &tags, query, postID)
	if err != nil {
		return nil, fmt.Errorf("ListTags: %w", err)
	}
	return tags, nil
}

func (r *tagRepository) ListAllTags(ctx context.Context) ([]models.Tag, error) {
	query := `
	SELECT id, name
	FROM tags
	`
	var tags []models.Tag
	err := r.db.SelectContext(ctx, &tags, query)
	if err != nil {
		return nil, fmt.Errorf("ListAllTags: %w", err)
	}
	return tags, nil
}
