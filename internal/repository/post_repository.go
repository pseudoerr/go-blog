package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pseudoerr/go-blog/internal/models"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) CreatePost(ctx context.Context, post models.Post) (models.Post, error) {

	query := ` 
	INSERT INTO posts (user_id, title, body, created_at, updated_at) 
	VALUES ($1, $2, $3, now(), now())
	RETURNING post_id, created_at, updated_at
	`

	var created_at, updated_at sql.NullTime
	err := r.db.QueryRowContext(ctx, query,
		post.UserID,
		post.Title,
		post.Body,
	).Scan(&post.UserID, &created_at, &updated_at)
	if err != nil {
		return models.Post{}, fmt.Errorf("CreatePost: %w", err)
	}
	post.CreatedAt = created_at.Time
	post.UpdatedAt = updated_at.Time

	return post, nil
}

func (r *postRepository) GetPostByID(ctx context.Context, postID int64) (models.Post, error) {
	query := `
		SELECT post_id, user_id, title, body, created_at, updated_at
		FROM posts
		WHERE post_id = $1
	`
	var post models.Post
	err := r.db.QueryRowContext(ctx, query, postID).Scan(
		&post.PostID,
		&post.UserID,
		&post.Title,
		&post.Body,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Post{}, nil
		}
		return models.Post{}, fmt.Errorf("GetPostByID: %w", err)
	}

	return post, nil
}

func (r *postRepository) ListPosts(ctx context.Context, userID int64, limit, offset int) ([]models.Post, error) {
	query := ` 
		SELECT  
			post_id,
			user_id, 
			title, 
			body, 
			created_at, 
			updated_at
		FROM posts
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return []models.Post{}, fmt.Errorf("ListPosts: %w", err)
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var p models.Post
		err := rows.Scan(
			&p.PostID,
			&p.UserID,
			&p.Body,
			&p.CreatedAt,
			&p.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("ListPosts scan error: %w", err)
		}
		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ListPosts rows error: %w", err)
	}

	return posts, nil
}

func (r *postRepository) UpdatePost(ctx context.Context, post models.Post) (models.Post, error) {
	query := `  
		UPDATE posts
		SET title = $1,
			body = $2,
			updated_at = now()
		WHERE post_id = $3
		RETURNING post_id, user_id, title, body, created_at, updated_at
	`
	var updatedPost models.Post

	err := r.db.QueryRowContext(ctx, query,
		post.Title,
		post.Body,
		post.PostID,
	).Scan(
		&updatedPost,
		&updatedPost.UserID,
		&updatedPost.Title,
		&updatedPost.Body,
		&updatedPost.CreatedAt,
		&updatedPost.UpdatedAt)

	if err != nil {
		return models.Post{}, fmt.Errorf("CreatePost: %w", err)
	}

	return updatedPost, nil
}

func (r *postRepository) DeletePost(ctx context.Context, postID int64) error {
	query := `
	DELETE FROM posts
    WHERE post_id = $1
	`

	res, err := r.db.ExecContext(ctx, query, postID)
	if err != nil {
		return fmt.Errorf("Delete post: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeletePost rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("DeletePost: no rows affected")
	}

	return nil
}

func (r *postRepository) ListPostsByTag(ctx context.Context, tag string, limit, offset int) ([]models.Post, error) {
	query := `
        SELECT p.post_id, p.user_id, p.title, p.body, p.created_at, p.updated_at
        FROM posts p
        JOIN post_tags pt ON p.post_id = pt.post_id
        JOIN tags t ON pt.tag_id = t.id
        WHERE t.name = $1
        ORDER BY p.created_at DESC
        LIMIT $2 OFFSET $3
    `

	rows, err := r.db.QueryContext(ctx, query, tag, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("ListPostsByTag: %w", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		err := rows.Scan(
			&p.PostID,
			&p.UserID,
			&p.Title,
			&p.Body,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("ListPostsByTag scan: %w", err)
		}
		posts = append(posts, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ListPostsByTag rows: %w", err)
	}

	return posts, nil
}
