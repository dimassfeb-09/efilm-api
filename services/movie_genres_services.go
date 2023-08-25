package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/helpers"
	"github.com/dimassfeb-09/efilm-api.git/repository"
)

type MovieGenreService interface {
	Save(ctx context.Context, r *web.MovieGenreModelRequestPost) error
	Delete(ctx context.Context, movieID int, genreID int) error
	FindByID(ctx context.Context, movieID int) (*web.MovieGenreModelResponse, error)
	FindGenreExists(ctx context.Context, genreID int) error
}

type MovieGenreServiceImpl struct {
	DB                   *sql.DB
	MovieGenreRepository repository.MovieGenreRepository
}

func NewMovieGenreService(DB *sql.DB, genreRepository repository.MovieGenreRepository) MovieGenreService {
	return &MovieGenreServiceImpl{DB: DB, MovieGenreRepository: genreRepository}
}

func (service *MovieGenreServiceImpl) Save(ctx context.Context, r *web.MovieGenreModelRequestPost) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	for _, genreID := range r.GenreIDS {
		fmt.Println(r.MovieID)
		err := service.MovieGenreRepository.Save(ctx, tx, r.MovieID, genreID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *MovieGenreServiceImpl) Delete(ctx context.Context, movieID int, genreID int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	err = service.FindGenreExists(ctx, genreID)
	if err != nil {
		return err
	}

	err = service.MovieGenreRepository.Delete(ctx, tx, movieID, genreID)
	if err != nil {
		return err
	}

	return nil
}

func (service *MovieGenreServiceImpl) FindByID(ctx context.Context, movieID int) (*web.MovieGenreModelResponse, error) {

	result, err := service.MovieGenreRepository.FindByID(ctx, service.DB, movieID)
	if err != nil {
		return nil, err
	}

	movieGenre := &web.MovieGenreModelResponse{
		Movie: web.Movie{
			MovieID:     result.Movie.ID,
			Title:       result.Movie.Title,
			ReleaseDate: result.Movie.ReleaseDate,
		},
	}

	for _, genre := range result.Genres {
		movieGenre.Genres = append(movieGenre.Genres, web.Genre{
			GenreID: genre.ID,
			Name:    genre.Name,
		})
	}

	return movieGenre, nil
}

func (service *MovieGenreServiceImpl) FindGenreExists(ctx context.Context, genreID int) error {
	return service.MovieGenreRepository.FindGenreExists(ctx, service.DB, genreID)
}
