package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
	"time"
)

type MovieDirectorRepository interface {
	Save(ctx context.Context, tx *sql.Tx, movieID, directorID int) error
	Delete(ctx context.Context, tx *sql.Tx, movieID int, directorID int) error
	FindByID(ctx context.Context, db *sql.DB, movieID int) (*domain.MovieDirector, error)
	FindDirectorAtMovie(ctx context.Context, db *sql.DB, movieID, directorID int) (exists bool, err error)
}

type MovieDirectorRepositoryaImpl struct {
}

func NewMovieDirectorRepository() MovieDirectorRepository {
	return &MovieDirectorRepositoryaImpl{}
}

func (repository *MovieDirectorRepositoryaImpl) Save(ctx context.Context, tx *sql.Tx, movieID, directorID int) error {
	query := "INSERT INTO movie_directors (movie_id, director_id) VALUES ($1, $2)"
	_, err := tx.ExecContext(ctx, query, movieID, directorID)
	if err != nil {
		return err
	}

	return nil
}

func (repository *MovieDirectorRepositoryaImpl) Delete(ctx context.Context, tx *sql.Tx, movieID int, directorID int) error {
	query := "DELETE FROM movie_directors WHERE movie_id = $1 AND  director_id = $2"
	_, err := tx.ExecContext(ctx, query, movieID, directorID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("data directors by ID at movie not found")
		}
		return errors.New("failed deleted directors at movie")
	}

	return nil
}

func (repository *MovieDirectorRepositoryaImpl) FindByID(ctx context.Context, db *sql.DB, movieID int) (*domain.MovieDirector, error) {

	query := `
		SELECT
			m.id AS movie_id,
			m.title AS title,
			m.release_date AS release_date,
			d.id AS director_id,
			d.name AS director_name,
			d.date_of_birth AS director_dob
		FROM movies m
		LEFT JOIN movie_directors md ON m.id = md.movie_id
		LEFT JOIN directors d ON md.director_id = d.id
		WHERE m.id = $1;
	`

	rows, err := db.QueryContext(ctx, query, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var directorMovie domain.MovieDirector
	for rows.Next() {
		var movieID int
		var title string
		var releaseDate time.Time
		var directorID sql.NullInt64
		var directorName sql.NullString
		var directorDOB sql.NullTime

		rows.Scan(
			&movieID,
			&title,
			&releaseDate,
			&directorID,
			&directorName,
			&directorDOB,
		)

		directorMovie.Movie = domain.Movie{
			ID:          movieID,
			Title:       title,
			ReleaseDate: releaseDate,
		}

		if directorID.Valid {
			director := domain.Director{
				ID:          int(directorID.Int64),
				Name:        directorName.String,
				DateOfBirth: directorDOB.Time,
			}
			directorMovie.Directors = append(directorMovie.Directors, director)
		} else {
			directorMovie.Directors = nil
		}

	}

	if directorMovie.Movie.ID == 0 {
		return nil, errors.New("movie not found")
	}

	return &directorMovie, nil
}

func (repository *MovieDirectorRepositoryaImpl) FindDirectorAtMovie(ctx context.Context, db *sql.DB, movieID, directorID int) (bool, error) {
	query := "SELECT movie_id FROM movie_directors WHERE movie_id = $1 AND director_id = $2"
	err := db.QueryRowContext(ctx, query, movieID, directorID).Scan(&movieID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errors.New("directors ID at movie not found")
		}
		fmt.Print(err)
		return false, err
	}
	return true, nil
}
