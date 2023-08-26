package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/helpers"
	"github.com/dimassfeb-09/efilm-api.git/repository"
)

type MovieDirectorService interface {
	Save(ctx context.Context, r *web.MovieDirectorModelRequestPost) error
	Delete(ctx context.Context, movieID int, directorID int) error
	FindByID(ctx context.Context, movieID int) (*web.MovieDirectorModelResponse, error)
	FindDirectorAtMovie(ctx context.Context, movieID, directorID int) (bool, error)
}

type MovieDirectorServiceImpl struct {
	DB                      *sql.DB
	MovieDirectorRepository repository.MovieDirectorRepository
	directorRepository      repository.DirectorRepository
	movieRepository         repository.MovieRepository
}

func NewMovieDirectorService(
	DB *sql.DB,
	movieDirectorRepository repository.MovieDirectorRepository,
) MovieDirectorService {
	return &MovieDirectorServiceImpl{
		DB:                      DB,
		MovieDirectorRepository: movieDirectorRepository,
		directorRepository:      repository.NewDirectorRepository(),
		movieRepository:         repository.NewMovieRepository(),
	}
}

func (service *MovieDirectorServiceImpl) Save(ctx context.Context, r *web.MovieDirectorModelRequestPost) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	_, err = service.directorRepository.FindByID(ctx, service.DB, r.DirectorID)
	if err != nil {
		return err
	}

	_, err = service.movieRepository.FindByID(ctx, service.DB, r.MovieID)
	if err != nil {
		return err
	}

	isExists, err := service.FindDirectorAtMovie(ctx, r.MovieID, r.DirectorID)
	// if director is on film, will response if director is already on film
	if isExists {
		return errors.New("the director is already on film")
	}

	return service.MovieDirectorRepository.Save(ctx, tx, r.MovieID, r.DirectorID)
}

func (service *MovieDirectorServiceImpl) Delete(ctx context.Context, movieID int, directorID int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	isExists, err := service.FindDirectorAtMovie(ctx, movieID, directorID)
	if err != nil && isExists == false {
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

func (service *MovieDirectorServiceImpl) FindDirectorAtMovie(ctx context.Context, movieID, directorID int) (bool, error) {
	return service.MovieDirectorRepository.FindDirectorAtMovie(ctx, service.DB, movieID, directorID)
}
