package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gizaw/09-library-api/internal/domain"
	"github.com/gizaw/09-library-api/internal/models"

	"github.com/jackc/pgx/v5/pgconn"
)

type PostgressUserRepository struct {
	db *sql.DB
}

func NewPostgressUserRepository(db *sql.DB ) *PostgressUserRepository{
	return &PostgressUserRepository{
		db: db,
	}
}

func (r *PostgressUserRepository) Create(cxt context.Context, params CreateUserParams)(*models.User, error){
	const query = `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`
	user := &models.User{
		Name: params.Name,
		Email: params.Email,
		Role: domain.RoleMember,
	}
	if err := r.db.QueryRowContext(cxt, query, params.Name, params.Email, params.PasswordHash).Scan(
		&user.ID, &user.CreatedAt, &user.UpdatedAt,
	); err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) && pgError.Code == "23505" && pgError.ConstraintName == "users_email_key"{
			return nil , domain.ErrDuplicateEmail
		}
		return nil, err
	}
	return user, nil
}

func (r *PostgressUserRepository)GetByEmail(cxt context.Context, email string)(*models.User, error) {
	const query = `
		SELECT
			id,
			name,
			email,
			password_hash,
			created_at,
			updated_at, 
			role
		FROM users
		WHERE email = $1
	`
	user := &models.User{}
	if err := r.db.QueryRowContext(cxt, query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt,&user.Role,
	); err != nil  {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return user , nil
}

func(r *PostgressUserRepository)GetById(cxt context.Context, id int64)(*models.User, error){
	const query = `
		SELECT
			id,
			name,
			email,
			password_hash,
			created_at,
			updated_at,
			role 
		FROM users
		WHERE id = $1
	`
	user := &models.User{}
	if err := r.db.QueryRowContext(cxt, query, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt,&user.Role,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows){
			return nil , domain.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}