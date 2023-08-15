package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/helpers"
	"github.com/dimassfeb-09/efilm-api.git/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, r *web.AuthModelRequest) error
	Login(ctx context.Context, r *web.AuthModelRequest) (token string, err error)
	findByUsername(ctx context.Context, name string) (*domain.Auth, error)
	findByID(ctx context.Context, ID int) (*domain.Auth, error)
}

type AuthServiceImpl struct {
	DB             *sql.DB
	AuthRepository repository.AuthRepository
}

func NewAuthService(DB *sql.DB, authRepository repository.AuthRepository) AuthService {
	return &AuthServiceImpl{DB: DB, AuthRepository: authRepository}
}

func (a *AuthServiceImpl) Register(ctx context.Context, r *web.AuthModelRequest) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	isExists, _ := a.findByUsername(ctx, r.Username)
	if isExists != nil {
		return errors.New("username already exists")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	return a.AuthRepository.Register(ctx, tx, &domain.Auth{
		Username: r.Username,
		Password: string(hashPassword),
	})
}

func (a *AuthServiceImpl) Login(ctx context.Context, r *web.AuthModelRequest) (string, error) {
	tx, err := a.DB.Begin()
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	result, err := a.AuthRepository.Login(ctx, tx, r.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(r.Password))
	if err != nil {
		return "", errors.New("terjadi kesalahan, email/password salah")
	}

	return helpers.GenerateTokenJWT(result.ID, result.Username, result.Role)
}

func (a *AuthServiceImpl) findByID(ctx context.Context, ID int) (*domain.Auth, error) {
	return a.AuthRepository.FindByID(ctx, a.DB, ID)
}

func (a *AuthServiceImpl) findByUsername(ctx context.Context, name string) (*domain.Auth, error) {
	return a.AuthRepository.FindByUsername(ctx, a.DB, name)
}
