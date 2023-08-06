package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/helpers"
	"github.com/dimassfeb-09/efilm-api.git/repository"
	"time"
)

type DirectorService interface {
	Save(ctx context.Context, r *web.DirectorModelRequest) error
	Update(ctx context.Context, r *web.DirectorModelRequest) error
	Delete(ctx context.Context, ID int) error
	FindByID(ctx context.Context, ID int) (*web.DirectorModelResponse, error)
	FindByName(ctx context.Context, name string) (*web.DirectorModelResponse, error)
	FindByNational(ctx context.Context, nationalityID int) ([]*web.DirectorModelResponse, error)
	FindAll(ctx context.Context) ([]*web.DirectorModelResponse, error)
}

type DirectorServiceImpl struct {
	DB                 *sql.DB
	DirectorRepository repository.DirectorRepository
}

func NewDirectorService(DB *sql.DB, directorRepository repository.DirectorRepository) DirectorService {
	return &DirectorServiceImpl{DB: DB, DirectorRepository: directorRepository}
}

func (a *DirectorServiceImpl) Save(ctx context.Context, r *web.DirectorModelRequest) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = a.FindByName(ctx, r.Name)
	if err == nil {
		return errors.New("director name already exists")
	}

	date, err := time.Parse(time.DateOnly, r.DateOfBirth)
	if err != nil {
		return errors.New("Incorrect date format yyyy-dd-mm")
	}

	return a.DirectorRepository.Save(ctx, tx, &domain.Director{
		Name:          r.Name,
		DateOfBirth:   date,
		NationalityID: r.NationalityID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	})

	return nil
}

func (a *DirectorServiceImpl) Update(ctx context.Context, r *web.DirectorModelRequest) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	date, err := time.Parse(time.DateOnly, r.DateOfBirth)
	if err != nil {
		return errors.New("incorrect format date: yyyy-mm-dd")
	}

	_, err = a.FindByID(ctx, r.ID)
	if err != nil {
		return err
	}

	return a.DirectorRepository.Update(ctx, tx, &domain.Director{
		ID:            r.ID,
		Name:          r.Name,
		DateOfBirth:   date,
		NationalityID: r.NationalityID,
		UpdatedAt:     time.Now(),
	})
}

func (a *DirectorServiceImpl) Delete(ctx context.Context, ID int) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = a.FindByID(ctx, ID)
	if err != nil {
		return err
	}

	err = a.DirectorRepository.Delete(ctx, tx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *DirectorServiceImpl) FindByID(ctx context.Context, ID int) (*web.DirectorModelResponse, error) {
	result, err := a.DirectorRepository.FindByID(ctx, a.DB, ID)
	if err != nil {
		return nil, err
	}

	return &web.DirectorModelResponse{
		ID:            result.ID,
		Name:          result.Name,
		DateOfBirth:   result.DateOfBirth,
		NationalityID: result.NationalityID,
		CreatedAt:     result.CreatedAt,
		UpdatedAt:     result.UpdatedAt,
	}, nil
}

func (a *DirectorServiceImpl) FindByName(ctx context.Context, name string) (*web.DirectorModelResponse, error) {

	result, err := a.DirectorRepository.FindByName(ctx, a.DB, name)
	if err != nil {
		return nil, err
	}

	return &web.DirectorModelResponse{
		ID:            result.ID,
		Name:          result.Name,
		DateOfBirth:   result.DateOfBirth,
		NationalityID: result.NationalityID,
		CreatedAt:     result.CreatedAt,
		UpdatedAt:     result.UpdatedAt,
	}, nil
}

func (a *DirectorServiceImpl) FindByNational(ctx context.Context, nationalityID int) ([]*web.DirectorModelResponse, error) {
	results, err := a.DirectorRepository.FindByNational(ctx, a.DB, nationalityID)
	if err != nil {
		return nil, err
	}

	var responses []*web.DirectorModelResponse
	for _, result := range results {
		response := web.DirectorModelResponse{
			ID:            result.ID,
			Name:          result.Name,
			DateOfBirth:   result.DateOfBirth,
			NationalityID: result.NationalityID,
			CreatedAt:     result.CreatedAt,
			UpdatedAt:     result.UpdatedAt,
		}
		responses = append(responses, &response)
	}

	return responses, nil
}

func (a *DirectorServiceImpl) FindAll(ctx context.Context) ([]*web.DirectorModelResponse, error) {
	results, err := a.DirectorRepository.FindAll(ctx, a.DB)
	if err != nil {
		return nil, err
	}

	var responses []*web.DirectorModelResponse
	for _, result := range results {
		response := web.DirectorModelResponse{
			ID:            result.ID,
			Name:          result.Name,
			DateOfBirth:   result.DateOfBirth,
			NationalityID: result.NationalityID,
			CreatedAt:     result.CreatedAt,
			UpdatedAt:     result.UpdatedAt,
		}

		responses = append(responses, &response)
	}

	return responses, nil
}
