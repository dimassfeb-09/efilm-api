package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/helpers"
	"github.com/dimassfeb-09/efilm-api.git/repository"
)

type GenreService interface {
	Save(ctx context.Context, r *web.GenreModelRequest) error
	Update(ctx context.Context, r *web.GenreModelRequest) error
	Delete(ctx context.Context, ID int) error
	FindAll(ctx context.Context) ([]*web.GenreModelResponse, error)
	FindByID(ctx context.Context, ID int) (*web.GenreModelResponse, error)
	FindByName(ctx context.Context, name string) (*web.GenreModelResponse, error)
}

type GenreServiceImpl struct {
	DB              *sql.DB
	GenreRepository repository.GenreRepository
}

func NewGenreService(DB *sql.DB, genreRepository repository.GenreRepository) GenreService {
	return &GenreServiceImpl{DB: DB, GenreRepository: genreRepository}
}

func (a *GenreServiceImpl) Save(ctx context.Context, r *web.GenreModelRequest) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = a.FindByName(ctx, r.Name)
	if err == nil {
		return errors.New("genre name already exists")
	}

	return a.GenreRepository.Save(ctx, tx, &domain.Genre{
		Name: r.Name,
	})
}

func (a *GenreServiceImpl) Update(ctx context.Context, r *web.GenreModelRequest) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = a.FindByID(ctx, r.ID)
	if err != nil {
		return err
	}

	return a.GenreRepository.Update(ctx, tx, &domain.Genre{
		ID:   r.ID,
		Name: r.Name,
	})
}

func (a *GenreServiceImpl) Delete(ctx context.Context, ID int) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = a.FindByID(ctx, ID)
	if err != nil {
		return err
	}

	err = a.GenreRepository.Delete(ctx, tx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *GenreServiceImpl) FindByID(ctx context.Context, ID int) (*web.GenreModelResponse, error) {
	result, err := a.GenreRepository.FindByID(ctx, a.DB, ID)
	if err != nil {
		return nil, err
	}

	return &web.GenreModelResponse{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (a *GenreServiceImpl) FindByName(ctx context.Context, name string) (*web.GenreModelResponse, error) {

	result, err := a.GenreRepository.FindByName(ctx, a.DB, name)
	if err != nil {
		return nil, err
	}

	return &web.GenreModelResponse{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (a *GenreServiceImpl) FindAll(ctx context.Context) ([]*web.GenreModelResponse, error) {
	results, err := a.GenreRepository.FindAll(ctx, a.DB)
	if err != nil {
		return nil, err
	}

	var responses []*web.GenreModelResponse
	for _, result := range results {
		response := web.GenreModelResponse{
			ID:        result.ID,
			Name:      result.Name,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}

		responses = append(responses, &response)
	}

	return responses, nil
}
