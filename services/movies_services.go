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

type MovieService interface {
	Save(ctx context.Context, r *web.MovieModelRequest) error
	Update(ctx context.Context, r *web.MovieModelRequest) error
	Delete(ctx context.Context, ID int) error
	FindByID(ctx context.Context, ID int) (*web.MovieModelResponse, error)
	FindByTitle(ctx context.Context, name string) (*web.MovieModelResponse, error)
	FindAll(ctx context.Context) ([]*web.MovieModelResponse, error)
}

type MovieServiceImpl struct {
	DB              *sql.DB
	MovieRepository repository.MovieRepository
}

func NewMovieService(DB *sql.DB, movieRepository repository.MovieRepository) MovieService {
	return &MovieServiceImpl{DB: DB, MovieRepository: movieRepository}
}

func (a *MovieServiceImpl) Save(ctx context.Context, r *web.MovieModelRequest) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = a.FindByTitle(ctx, r.Title)
	if err == nil {
		return errors.New("movie title already exists")
	}

	releaseDate, err := time.Parse(time.DateOnly, r.ReleaseDate)
	if err != nil {
		return errors.New("incorrect date format yyyy-dd-mm")
	}
	return a.MovieRepository.Save(ctx, tx, &domain.Movie{
		Title:       r.Title,
		ReleaseDate: releaseDate,
		Duration:    r.Duration,
		Plot:        r.Plot,
		PosterUrl:   r.PosterUrl,
		TrailerUrl:  r.TrailerUrl,
		Language:    r.Language,
	})

	return nil
}

func (a *MovieServiceImpl) Update(ctx context.Context, r *web.MovieModelRequest) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = a.FindByID(ctx, r.ID)
	if err != nil {
		return err
	}

	releaseDate, err := time.Parse(time.DateOnly, r.ReleaseDate)
	if err != nil {
		return errors.New("incorrect date format yyyy-dd-mm")
	}

	return a.MovieRepository.Update(ctx, tx, &domain.Movie{
		ID:          r.ID,
		Title:       r.Title,
		ReleaseDate: releaseDate,
		Duration:    r.Duration,
		Plot:        r.Plot,
		PosterUrl:   r.PosterUrl,
		TrailerUrl:  r.TrailerUrl,
		Language:    r.Language,
	})
}

func (a *MovieServiceImpl) Delete(ctx context.Context, ID int) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = a.FindByID(ctx, ID)
	if err != nil {
		return err
	}

	err = a.MovieRepository.Delete(ctx, tx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *MovieServiceImpl) FindByID(ctx context.Context, ID int) (*web.MovieModelResponse, error) {
	result, err := a.MovieRepository.FindByID(ctx, a.DB, ID)
	if err != nil {
		return nil, err
	}

	return &web.MovieModelResponse{
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
	}, nil
}

func (a *MovieServiceImpl) FindByTitle(ctx context.Context, name string) (*web.MovieModelResponse, error) {

	result, err := a.MovieRepository.FindByTitle(ctx, a.DB, name)
	if err != nil {
		return nil, err
	}

	return &web.MovieModelResponse{
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
	}, nil
}

func (a *MovieServiceImpl) FindAll(ctx context.Context) ([]*web.MovieModelResponse, error) {
	results, err := a.MovieRepository.FindAll(ctx, a.DB)
	if err != nil {
		return nil, err
	}

	var responses []*web.MovieModelResponse
	for _, result := range results {
		response := web.MovieModelResponse{
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
		}

		responses = append(responses, &response)
	}

	return responses, nil
}
