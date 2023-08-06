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

type NationalService interface {
	Save(ctx context.Context, r *web.NationalModelRequest) error
	Update(ctx context.Context, r *web.NationalModelRequest) error
	Delete(ctx context.Context, ID int) error
	FindAll(ctx context.Context) ([]*web.NationalModelResponse, error)
	FindByID(ctx context.Context, ID int) (*web.NationalModelResponse, error)
	FindByName(ctx context.Context, name string) (*web.NationalModelResponse, error)
}

type NationalServiceImpl struct {
	DB                 *sql.DB
	NationalRepository repository.NationalRepository
}

func NewNationalService(DB *sql.DB, nationalRepository repository.NationalRepository) NationalService {
	return &NationalServiceImpl{DB: DB, NationalRepository: nationalRepository}
}

func (a *NationalServiceImpl) Save(ctx context.Context, r *web.NationalModelRequest) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, errs := a.FindByName(ctx, r.Name)
	if errs == nil {
		return errors.New("national name already exists")
	}

	return a.NationalRepository.Save(ctx, tx, &domain.National{
		Name: r.Name,
	})
}

func (a *NationalServiceImpl) Update(ctx context.Context, r *web.NationalModelRequest) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = a.FindByID(ctx, r.ID)
	if err != nil {
		return err
	}

	return a.NationalRepository.Update(ctx, tx, &domain.National{
		ID:   r.ID,
		Name: r.Name,
	})
}

func (a *NationalServiceImpl) Delete(ctx context.Context, ID int) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = a.FindByID(ctx, ID)
	if err != nil {
		return err
	}

	err = a.NationalRepository.Delete(ctx, tx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *NationalServiceImpl) FindByID(ctx context.Context, ID int) (*web.NationalModelResponse, error) {
	result, err := a.NationalRepository.FindByID(ctx, a.DB, ID)
	if err != nil {
		return nil, err
	}

	return &web.NationalModelResponse{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (a *NationalServiceImpl) FindByName(ctx context.Context, name string) (*web.NationalModelResponse, error) {

	result, err := a.NationalRepository.FindByName(ctx, a.DB, name)
	if err != nil {
		return nil, err
	}

	return &web.NationalModelResponse{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (a *NationalServiceImpl) FindAll(ctx context.Context) ([]*web.NationalModelResponse, error) {
	results, err := a.NationalRepository.FindAll(ctx, a.DB)
	if err != nil {
		return nil, err
	}

	var responses []*web.NationalModelResponse
	for _, result := range results {
		response := web.NationalModelResponse{
			ID:        result.ID,
			Name:      result.Name,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}

		responses = append(responses, &response)
	}

	return responses, nil
}
