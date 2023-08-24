package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
	"time"
)

type MovieGenreRepository interface {
	Save(ctx context.Context, tx *sql.Tx, movieID, genreID int) error
	Delete(ctx context.Context, tx *sql.Tx, movieID int, genreID int) error
	FindByID(ctx context.Context, db *sql.DB, movieID int) (*domain.MovieGenre, error)
	FindGenreExists(ctx context.Context, db *sql.DB, genreID int) error
}

type MovieGenreRepositoryaImpl struct {
}

func NewMovieGenreRepository() MovieGenreRepository {
	return &MovieGenreRepositoryaImpl{}
}

func (repository *MovieGenreRepositoryaImpl) Save(ctx context.Context, tx *sql.Tx, movieID, genreID int) error {
	query := "INSERT INTO movie_genres (movie_id, genre_id) VALUES ($1, $2)"
	_, err := tx.ExecContext(ctx, query, movieID, genreID)
	if err != nil {
		return err
	}

	return nil
}

func (repository *MovieGenreRepositoryaImpl) Delete(ctx context.Context, tx *sql.Tx, movieID int, genreID int) error {
	query := "DELETE FROM movie_genres WHERE movie_id = $1 AND  genre_id = $2"
	_, err := tx.ExecContext(ctx, query, movieID, genreID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("data genres by ID at movie not found")
		}
		return errors.New("failed deleted genres at movie")
	}

	return nil
}

func (repository *MovieGenreRepositoryaImpl) FindByID(ctx context.Context, db *sql.DB, movieID int) (*domain.MovieGenre, error) {

	query := `
		SELECT
			m.id AS movie_id,
			m.title AS title,
			m.release_date AS release_date,
			g.id AS genre_id,
			g.name AS genre_name
		FROM movies m
		LEFT JOIN movie_genres mg ON m.id = mg.movie_id
		LEFT JOIN genres g ON mg.genre_id = g.id
		WHERE m.id = $1;
	`

	rows, err := db.QueryContext(ctx, query, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genreMovie domain.MovieGenre
	for rows.Next() {
		var movieID int
		var title string
		var releaseDate time.Time
		var genreID sql.NullInt64
		var genreName sql.NullString

		rows.Scan(
			&movieID,
			&title,
			&releaseDate,
			&genreID,
			&genreName,
		)

		genreMovie.Movie = domain.Movie{
			ID:          movieID,
			Title:       title,
			ReleaseDate: releaseDate,
		}

		if genreID.Valid {
			genreMovie.GenreIDS = append(genreMovie.GenreIDS, int(genreID.Int64))
		} else {
			genreMovie.GenreIDS = nil
		}

	}

	if genreMovie.Movie.ID == 0 {
		return nil, errors.New("movie not found")
	}

	return &genreMovie, nil
}

func (repository *MovieGenreRepositoryaImpl) FindGenreExists(ctx context.Context, db *sql.DB, genreID int) error {
	query := "SELECT genre_id FROM movie_genres WHERE genre_id = $1"
	err := db.QueryRowContext(ctx, query, genreID).Scan(&genreID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("genres ID at movie not found")
		}
		return err
	}
	return nil
}
