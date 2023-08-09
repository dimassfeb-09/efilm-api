package services

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/helpers"
	"github.com/dimassfeb-09/efilm-api.git/repository"
)

type MovieDirectorService interface {
	Save(ctx context.Context, r *web.MovieDirectorModelRequestPost) error
	Delete(ctx context.Context, movieID int, directorID int) error
	FindByID(ctx context.Context, movieID int) (*web.MovieDirectorModelResponse, error)
	FindDirectorExists(ctx context.Context, directorID int) error
}

type MovieDirectorServiceImpl struct {
	DB                      *sql.DB
	MovieDirectorRepository repository.MovieDirectorRepository
}

func NewMovieDirectorService(DB *sql.DB, directorRepository repository.MovieDirectorRepository) MovieDirectorService {
	return &MovieDirectorServiceImpl{DB: DB, MovieDirectorRepository: directorRepository}
}

func (service *MovieDirectorServiceImpl) Save(ctx context.Context, r *web.MovieDirectorModelRequestPost) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	err = service.MovieDirectorRepository.Save(ctx, tx, r.MovieID, r.DirectorID)
	if err != nil {
		return err
	}

	return nil
}

func (service *MovieDirectorServiceImpl) Delete(ctx context.Context, movieID int, directorID int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	err = service.FindDirectorExists(ctx, directorID)
	if err != nil {
		return err
	}

	err = service.MovieDirectorRepository.Delete(ctx, tx, movieID, directorID)
	if err != nil {
		return err
	}

	return nil
}

func (service *MovieDirectorServiceImpl) FindByID(ctx context.Context, movieID int) (*web.MovieDirectorModelResponse, error) {

	result, err := service.MovieDirectorRepository.FindByID(ctx, service.DB, movieID)
	if err != nil {
		return nil, err
	}

	movieDirector := &web.MovieDirectorModelResponse{
		Movie: web.Movie{
			MovieID:     result.Movie.ID,
			Title:       result.Movie.Title,
			ReleaseDate: result.Movie.ReleaseDate,
		},
	}

	for _, director := range result.Directors {
		movieDirector.Directors = append(movieDirector.Directors, web.Director{
			DirectorID:  director.ID,
			Name:        director.Name,
			DateOfBirth: director.DateOfBirth,
		})
	}

	return movieDirector, nil
}

func (service *MovieDirectorServiceImpl) FindDirectorExists(ctx context.Context, directorID int) error {
	return service.MovieDirectorRepository.FindDirectorExists(ctx, service.DB, directorID)
}
