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
	FindAllMoviesByID(ctx context.Context, ID int) (*web.MoviesGenreResponse, error)
	FindByID(ctx context.Context, ID int) (*web.GenreModelResponse, error)
	FindByName(ctx context.Context, name string) (*web.GenreModelResponse, error)
}

type GenreServiceImpl struct {
	DB              *sql.DB
	GenreRepository repository.GenreRepository
	MovieService    MovieService
}

func NewGenreService(DB *sql.DB, genreRepository repository.GenreRepository, movieService MovieService) GenreService {
	return &GenreServiceImpl{DB: DB, GenreRepository: genreRepository, MovieService: movieService}
}

func (service *GenreServiceImpl) Save(ctx context.Context, r *web.GenreModelRequest) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = service.FindByName(ctx, r.Name)
	if err == nil {
		return errors.New("genre name already exists")
	}

	return service.GenreRepository.Save(ctx, tx, &domain.Genre{
		Name: r.Name,
	})
}

func (service *GenreServiceImpl) Update(ctx context.Context, r *web.GenreModelRequest) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = service.FindByID(ctx, r.ID)
	if err != nil {
		return err
	}

	return service.GenreRepository.Update(ctx, tx, &domain.Genre{
		ID:   r.ID,
		Name: r.Name,
	})
}

func (service *GenreServiceImpl) Delete(ctx context.Context, ID int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = service.FindByID(ctx, ID)
	if err != nil {
		return err
	}

	err = service.GenreRepository.Delete(ctx, tx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (service *GenreServiceImpl) FindByID(ctx context.Context, ID int) (*web.GenreModelResponse, error) {
	result, err := service.GenreRepository.FindByID(ctx, service.DB, ID)
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

func (service *GenreServiceImpl) FindByName(ctx context.Context, name string) (*web.GenreModelResponse, error) {

	result, err := service.GenreRepository.FindByName(ctx, service.DB, name)
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

func (service *GenreServiceImpl) FindAll(ctx context.Context) ([]*web.GenreModelResponse, error) {
	results, err := service.GenreRepository.FindAll(ctx, service.DB)
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

func (service *GenreServiceImpl) FindAllMoviesByID(ctx context.Context, ID int) (*web.MoviesGenreResponse, error) {
	_, err := service.GenreRepository.FindByID(ctx, service.DB, ID)
	if err != nil {
		return nil, err
	}

	results, err := service.MovieService.FindAllMoviesByGenreID(ctx, ID)
	if err != nil {
		return nil, err
	}

	var responses web.MoviesGenreResponse
	for _, result := range results {
		responses.Movies = append(responses.Movies, &web.MovieModelResponse{
			ID:          result.ID,
			Title:       result.Title,
			ReleaseDate: result.ReleaseDate,
			Duration:    result.Duration,
			Plot:        result.Plot,
			PosterUrl:   result.PosterUrl,
			TrailerUrl:  result.TrailerUrl,
			Language:    result.Language,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
		})
	}
	responses.GenreID = ID

	return &responses, nil
}
