package repository

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pseudoerr/go-blog/internal/models"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, username string, hashedPassword string) (models.User, error) {
	userID := uuid.Must(uuid.NewV4())

	query := `INSERT INTO users (user_id, username, hashed_password) VALUES ($1, $2, $3) RETURNING user_id, username, hashed_password, time_stamp`

	var user models.User
	err := r.db.QueryRowContext(ctx, query, userID, username, hashedPassword).Scan(
		&user.UserID,
		&user.Username,
		&user.HashedPassword,
		&user.TimeStamp,
	)

	return user, err
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	query := `SELECT user_id, username, hashed_password, time_stamp FROM users WHERE username = $1`

	var user models.User
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.UserID,
		&user.Username,
		&user.HashedPassword,
		&user.TimeStamp,
	)

	return user, err
}

func (r *userRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	query := `SELECT user_id, username, hashed_password, time_stamp FROM users WHERE user_id = $1`

	var user models.User
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&user.UserID,
		&user.Username,
		&user.HashedPassword,
		&user.TimeStamp,
	)

	return user, err
}
