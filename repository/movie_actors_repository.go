package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
	"time"
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
			m.id AS movie_id,
			m.title AS title,
			m.release_date AS release_date,
			a.id AS actor_id,
			a.name AS actor_name,
			a.date_of_birth AS actor_dob
		FROM movies m
		LEFT JOIN movie_actors ma ON m.id = ma.movie_id
		LEFT JOIN actors a ON ma.actor_id = a.id
		WHERE m.id = $1;
	`

	rows, err := db.QueryContext(ctx, query, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actorMovie domain.MovieActor
	for rows.Next() {
		var movieID int
		var title string
		var releaseDate time.Time
		var actorID sql.NullInt64
		var actorName sql.NullString
		var actorDOB sql.NullTime

		rows.Scan(
			&movieID,
			&title,
			&releaseDate,
			&actorID,
			&actorName,
			&actorDOB,
		)

		actorMovie.Movie = domain.Movie{
			ID:          movieID,
			Title:       title,
			ReleaseDate: releaseDate,
		}

		if actorID.Valid {
			actor := domain.Actor{
				ID:          int(actorID.Int64),
				Name:        actorName.String,
				DateOfBirth: actorDOB.Time,
			}
			actorMovie.Actors = append(actorMovie.Actors, actor)
		} else {
			actorMovie.Actors = nil
		}

	}

	if actorMovie.Movie.ID == 0 {
		return nil, errors.New("movie not found")
	}

	return &actorMovie, nil
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
