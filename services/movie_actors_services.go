package services

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/efilm-api.git/entity/web"
	"github.com/dimassfeb-09/efilm-api.git/helpers"
	"github.com/dimassfeb-09/efilm-api.git/repository"
)

type MovieActorService interface {
	Save(ctx context.Context, r *web.MovieActorModelRequestPost) error
	Update(ctx context.Context, r *web.MovieActorModelRequestPut) error
	Delete(ctx context.Context, actorID int) error
	FindByID(ctx context.Context, movieID int) (*web.MovieActorModelResponse, error)
}

type MovieActorServiceImpl struct {
	DB                   *sql.DB
	MovieActorRepository repository.MovieActorRepository
}

func NewMovieActorService(DB *sql.DB, actorRepository repository.MovieActorRepository) MovieActorService {
	return &MovieActorServiceImpl{DB: DB, MovieActorRepository: actorRepository}
}

func (service *MovieActorServiceImpl) Save(ctx context.Context, r *web.MovieActorModelRequestPost) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	err = service.MovieActorRepository.Save(ctx, tx, r.MovieID, r.ActorID, r.Role)
	if err != nil {
		return err
	}

	return nil
}

func (service *MovieActorServiceImpl) Update(ctx context.Context, r *web.MovieActorModelRequestPut) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	err = service.MovieActorRepository.FindActorAtMovieExists(ctx, service.DB, r.ActorID)
	if err != nil {
		return err
	}

	err = service.MovieActorRepository.Update(ctx, tx, r.MovieID, r.ActorID, r.Role)
	if err != nil {
		return err
	}

	return nil
}

func (service *MovieActorServiceImpl) Delete(ctx context.Context, actorID int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}
	defer helpers.RollbackOrCommit(ctx, tx)

	err = service.MovieActorRepository.FindActorAtMovieExists(ctx, service.DB, actorID)
	if err != nil {
		return err
	}

	err = service.MovieActorRepository.Delete(ctx, tx, actorID)
	if err != nil {
		return err
	}

	return nil
}

func (service *MovieActorServiceImpl) FindByID(ctx context.Context, movieID int) (*web.MovieActorModelResponse, error) {
	result, err := service.MovieActorRepository.FindByID(ctx, service.DB, movieID)
	if err != nil {
		return nil, err
	}

	movieActor := &web.MovieActorModelResponse{
		Movie: web.Movie{
			MovieID:     result.Movie.ID,
			Title:       result.Movie.Title,
			ReleaseDate: result.Movie.ReleaseDate,
		},
	}

	for _, actor := range result.Actors {
		movieActor.Actors = append(movieActor.Actors, web.Actor{
			ActorID:     actor.ID,
			Name:        actor.Name,
			DateOfBirth: actor.DateOfBirth,
			Role:        actor.Role,
		})
	}

	return movieActor, nil
}
