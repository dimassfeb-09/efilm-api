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

type ActorService interface {
	Save(ctx context.Context, r *web.ActorModelRequest) error
	Update(ctx context.Context, r *web.ActorModelRequest) error
	Delete(ctx context.Context, ID int) error
	FindByID(ctx context.Context, ID int) (*web.ActorModelResponse, error)
	FindByName(ctx context.Context, name string) (*web.ActorModelResponse, error)
	FindByNational(ctx context.Context, nationalityID int) ([]*web.ActorModelResponse, error)
	FindAll(ctx context.Context) ([]*web.ActorModelResponse, error)
}

type ActorServiceImpl struct {
	DB              *sql.DB
	ActorRepository repository.ActorRepository
}

func NewActorService(DB *sql.DB, actorRepository repository.ActorRepository) ActorService {
	return &ActorServiceImpl{DB: DB, ActorRepository: actorRepository}
}

func (a *ActorServiceImpl) Save(ctx context.Context, r *web.ActorModelRequest) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = a.FindByName(ctx, r.Name)
	if err == nil {
		return errors.New("actors name already exists")
	}

	date, err := time.Parse(time.DateOnly, r.DateOfBirth)
	if err != nil {
		return errors.New("Incorrect date format yyyy-dd-mm")
	}

	return a.ActorRepository.Save(ctx, tx, &domain.Actor{
		Name:          r.Name,
		DateOfBirth:   date,
		NationalityID: r.NationalityID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	})

	return nil
}

func (a *ActorServiceImpl) Update(ctx context.Context, r *web.ActorModelRequest) error {
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

	return a.ActorRepository.Update(ctx, tx, &domain.Actor{
		ID:            r.ID,
		Name:          r.Name,
		DateOfBirth:   date,
		NationalityID: r.NationalityID,
		UpdatedAt:     time.Now(),
	})
}

func (a *ActorServiceImpl) Delete(ctx context.Context, ID int) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = a.FindByID(ctx, ID)
	if err != nil {
		return err
	}

	err = a.ActorRepository.Delete(ctx, tx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *ActorServiceImpl) FindByID(ctx context.Context, ID int) (*web.ActorModelResponse, error) {
	result, err := a.ActorRepository.FindByID(ctx, a.DB, ID)
	if err != nil {
		return nil, err
	}

	return &web.ActorModelResponse{
		ID:            result.ID,
		Name:          result.Name,
		DateOfBirth:   result.DateOfBirth,
		NationalityID: result.NationalityID,
		CreatedAt:     result.CreatedAt,
		UpdatedAt:     result.UpdatedAt,
	}, nil
}

func (a *ActorServiceImpl) FindByName(ctx context.Context, name string) (*web.ActorModelResponse, error) {

	result, err := a.ActorRepository.FindByName(ctx, a.DB, name)
	if err != nil {
		return nil, err
	}

	return &web.ActorModelResponse{
		ID:            result.ID,
		Name:          result.Name,
		DateOfBirth:   result.DateOfBirth,
		NationalityID: result.NationalityID,
		CreatedAt:     result.CreatedAt,
		UpdatedAt:     result.UpdatedAt,
	}, nil
}

func (a *ActorServiceImpl) FindByNational(ctx context.Context, nationalityID int) ([]*web.ActorModelResponse, error) {
	results, err := a.ActorRepository.FindByNational(ctx, a.DB, nationalityID)
	if err != nil {
		return nil, err
	}

	var responses []*web.ActorModelResponse
	for _, result := range results {
		response := web.ActorModelResponse{
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

func (a *ActorServiceImpl) FindAll(ctx context.Context) ([]*web.ActorModelResponse, error) {
	results, err := a.ActorRepository.FindAll(ctx, a.DB)
	if err != nil {
		return nil, err
	}

	var responses []*web.ActorModelResponse
	for _, result := range results {
		response := web.ActorModelResponse{
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
