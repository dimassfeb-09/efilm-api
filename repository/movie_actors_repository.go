package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
)

type MovieActorRepository interface {
	Save(ctx context.Context, tx *sql.Tx, movieID, actorID int, role string) error
	Update(ctx context.Context, tx *sql.Tx, movieID, actorID int, role string) error
	Delete(ctx context.Context, tx *sql.Tx, actorID int) error
	FindByID(ctx context.Context, db *sql.DB, movieID int) (*domain.MovieActor, error)
	FindActorAtMovieExists(ctx context.Context, db *sql.DB, actorID int) error
}

type MovieActorRepositoryaImpl struct {
}

func NewMovieActorRepository() MovieActorRepository {
	return &MovieActorRepositoryaImpl{}
}

func (repository *MovieActorRepositoryaImpl) Save(ctx context.Context, tx *sql.Tx, movieID, actorID int, role string) error {
	query := "INSERT INTO movie_actors (movie_id, actor_id, role) VALUES ($1, $2, $3)"
	_, err := tx.ExecContext(ctx, query, movieID, actorID, role)
	if err != nil {
		return err
	}

	return nil
}

func (repository *MovieActorRepositoryaImpl) Update(ctx context.Context, tx *sql.Tx, movieID, actorID int, role string) error {
	query := "UPDATE movie_actors SET movie_id = $1, actor_id = $2, role = $3 WHERE actor_id = $4"
	_, err := tx.ExecContext(ctx, query, movieID, actorID, role, actorID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("data actors by ID at movie not found")
		}
	}

	return nil
}

func (repository *MovieActorRepositoryaImpl) Delete(ctx context.Context, tx *sql.Tx, actorID int) error {
	query := "DELETE FROM movie_actors WHERE actor_id = $1"
	_, err := tx.ExecContext(ctx, query, actorID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("data actors by ID at movie not found")
		}
		return err
	}

	return nil
}

func (repository *MovieActorRepositoryaImpl) FindByID(ctx context.Context, db *sql.DB, movieID int) (*domain.MovieActor, error) {

	query := `
		SELECT 
		    movie_actors.id AS id,
		    movies.id AS movie_id,
			movies.title AS title,
			movies.release_date AS release_date,
			movies.duration AS duration,
			movies.plot AS plot,
			movies.poster_url AS poster_url,
			movies.trailer_url AS trailer_url,
			movies.language AS language,
			movies.created_at AS movie_created_at,
			movies.updated_at AS movie_updated_at,
			actors.id AS actor_id,
			actors.name AS name,
			actors.date_of_birth AS date_of_birth,
			actors.nationality_id AS nationality_id,
			actors.created_at AS actor_created_at,
			actors.updated_at AS actor_updated_at
		FROM movie_actors
		JOIN actors ON movie_actors.actor_id = actors.id
		JOIN movies ON movie_actors.movie_id = movies.id WHERE movies.id = $1
	`

	rows, err := db.QueryContext(ctx, query, movieID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed get data from database")
		}
		return nil, err
	}
	defer rows.Close()

	var movieActors domain.MovieActor
	var movie domain.Movie
	for rows.Next() {
		var actor domain.Actor

		err := rows.Scan(
			&movieActors.ID,
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.Duration,
			&movie.Plot,
			&movie.PosterUrl,
			&movie.TrailerUrl,
			&movie.Language,
			&movie.CreatedAt,
			&movie.UpdatedAt,
			&actor.ID,
			&actor.Name,
			&actor.DateOfBirth,
			&actor.NationalityID,
			&actor.CreatedAt,
			&actor.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		movieActors.MovieID = movie.ID
		movieActors.Title = movie.Title
		movieActors.ReleaseDate = movie.ReleaseDate
		movieActors.Duration = movie.Duration
		movieActors.Plot = movie.Plot
		movieActors.PosterUrl = movie.PosterUrl
		movieActors.TrailerUrl = movie.TrailerUrl
		movieActors.Language = movie.Language
		movieActors.CreatedAt = movie.CreatedAt
		movieActors.UpdatedAt = movie.UpdatedAt

		actorInstance := domain.Actor{
			ID:            actor.ID,
			Name:          actor.Name,
			DateOfBirth:   actor.DateOfBirth,
			NationalityID: actor.NationalityID,
			CreatedAt:     actor.CreatedAt,
			UpdatedAt:     actor.UpdatedAt,
		}

		movieActors.Actors = append(movieActors.Actors, actorInstance)
	}

	return &movieActors, nil
}

func (repository *MovieActorRepositoryaImpl) FindActorAtMovieExists(ctx context.Context, db *sql.DB, actorID int) error {
	query := "SELECT id FROM movie_actors WHERE actor_id = $1"
	err := db.QueryRowContext(ctx, query, actorID).Scan()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("actors ID at movie not found")
		}
		return err
	}
	return nil
}
