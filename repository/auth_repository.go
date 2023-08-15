package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
)

type AuthRepository interface {
	Register(ctx context.Context, tx *sql.Tx, auth *domain.Auth) error
	Login(ctx context.Context, tx *sql.Tx, username string) (*domain.Auth, error)
	FindByUsername(ctx context.Context, db *sql.DB, username string) (*domain.Auth, error)
	FindByID(ctx context.Context, db *sql.DB, ID int) (*domain.Auth, error)
}

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (a *AuthRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, auth *domain.Auth) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := tx.ExecContext(ctx, query, auth.Username, auth.Password)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, username string) (*domain.Auth, error) {
	query := "SELECT id, username, password, role FROM users WHERE username = $1"

	var auth domain.Auth
	err := tx.QueryRowContext(ctx, query, username).
		Scan(&auth.ID, &auth.Username, &auth.Password, &auth.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("sorry, username not found")
		}
		return nil, err
	}

	return &auth, nil
}

func (a *AuthRepositoryImpl) FindByUsername(ctx context.Context, db *sql.DB, username string) (*domain.Auth, error) {
	var auth domain.Auth
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&auth.ID, &auth.Username, &auth.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("auth with username %s not found", username)
		}
		return nil, err
	}

	return &auth, nil
}

func (a *AuthRepositoryImpl) FindByID(ctx context.Context, db *sql.DB, ID int) (*domain.Auth, error) {
	var auth domain.Auth
	err := db.QueryRow("SELECT id FROM users WHERE id = $1", ID).Scan(&auth.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("auth with username %s not found", ID)
		}
		return nil, err
	}

	return &auth, nil
}
