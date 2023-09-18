package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
)

type MovieRepository interface {
	Save(ctx context.Context, tx *sql.Tx, movie *domain.Movie) (movieID int, err error)
	Update(ctx context.Context, tx *sql.Tx, movie *domain.Movie) error
	Delete(ctx context.Context, tx *sql.Tx, ID int) error
	FindByID(ctx context.Context, db *sql.DB, ID int) (*domain.Movie, error)
	FindByTitle(ctx context.Context, db *sql.DB, name string) (*domain.Movie, error)
	FindAll(ctx context.Context, db *sql.DB) ([]*domain.Movie, error)
	FindAllMoviesByGenreID(ctx context.Context, db *sql.DB, genreID int) ([]*domain.Movie, error)
}

type MovieRepositoryImpl struct {
}

func NewMovieRepository() MovieRepository {
	return &MovieRepositoryImpl{}
}

func (a *MovieRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, movie *domain.Movie) (int, error) {
	var id int
	query := `
		INSERT INTO 
		    movies (title, release_date, duration, plot, poster_url, trailer_url, language, nationality_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err := tx.QueryRowContext(ctx, query, movie.Title, movie.ReleaseDate, movie.Duration, movie.Plot, movie.PosterUrl, movie.TrailerUrl, movie.Language, movie.NationalID).
		Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (a *MovieRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, movie *domain.Movie) error {
	query := "UPDATE movies SET id = $1, title = $2, release_date = $3, duration = $4, plot = $5, poster_url = $6, trailer_url = $7, language = $8, nationality_id = $9, updated_at = CURRENT_TIMESTAMP WHERE id = $10"
	_, err := tx.ExecContext(ctx, query, movie.ID, movie.Title, movie.ReleaseDate, movie.Duration, movie.Plot, movie.PosterUrl, movie.TrailerUrl, movie.Language, movie.NationalID, movie.ID)
	if err != nil {
		fmt.Println(err)
		return errors.New("failed update data movie")
	}

	return nil
}

func (a *MovieRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, ID int) error {
	_, err := tx.Exec("DELETE FROM movies WHERE id = $1", ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *MovieRepositoryImpl) FindByID(ctx context.Context, db *sql.DB, ID int) (*domain.Movie, error) {

	query := `
		SELECT 
		    m.id as id, 
		    title, 
		    release_date, 
		    duration, 
		    plot, 
		    poster_url, 
		    trailer_url, 
		    language, 
		    n.id as national_id,
		    m.created_at as created_at, 
		    m.updated_at as updated_at 
		FROM movies as m
		JOIN national as n ON m.nationality_id = n.id
		WHERE m.id = $1`

	var movie domain.Movie
	row := db.QueryRowContext(ctx, query, ID)

	err := row.Scan(&movie.ID, &movie.Title, &movie.ReleaseDate, &movie.Duration, &movie.Plot, &movie.PosterUrl, &movie.TrailerUrl, &movie.Language, &movie.NationalID, &movie.CreatedAt, &movie.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("sorry, movie id not found")
		}
		return nil, err
	}

	return &movie, nil
}

func (a *MovieRepositoryImpl) FindByTitle(ctx context.Context, db *sql.DB, title string) (*domain.Movie, error) {
	var movie domain.Movie
	err := db.QueryRow("SELECT * FROM movies WHERE title = $1", title).
		Scan(&movie.ID, &movie.Title, &movie.ReleaseDate, &movie.Duration, &movie.Plot, &movie.PosterUrl, &movie.TrailerUrl, &movie.Language, &movie.CreatedAt, &movie.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("movie with name %s not found", title)
		}
		return nil, err
	}

	return &movie, nil
}

func (a *MovieRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) ([]*domain.Movie, error) {
	query := `
		SELECT 
		    m.id as id, 
		    title, 
		    release_date, 
		    duration, 
		    plot, 
		    poster_url, 
		    trailer_url, 
		    language, 
		    n.id as national_id,
		    m.created_at as created_at, 
		    m.updated_at as updated_at 
		FROM movies as m
		JOIN national as n ON m.nationality_id = n.id`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed get data from database")
		}
		return nil, err
	}

	var movies []*domain.Movie
	for rows.Next() {
		var movie domain.Movie
		rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.Duration,
			&movie.Plot,
			&movie.PosterUrl,
			&movie.TrailerUrl,
			&movie.Language,
			&movie.NationalID,
			&movie.CreatedAt,
			&movie.UpdatedAt)
		movies = append(movies, &movie)
	}

	return movies, nil
}

func (a *MovieRepositoryImpl) FindAllMoviesByGenreID(ctx context.Context, db *sql.DB, genreID int) ([]*domain.Movie, error) {
	query := `SELECT movies.id as movie_id,
		       movies.title as title,
		       movies.release_date as release_date,
		       movies.duration as duration,
		       movies.plot as plot,
		       movies.poster_url as poster_url,
		       movies.trailer_url as trailer_url,
		       movies.language as language,
		       movies.created_at as created_at,
		       movies.updated_at as updated_at,
		       movie_genres.genre_id as genre_id
		FROM movies 
		    JOIN movie_genres ON movies.id = movie_genres.movie_id 
		    JOIN genres ON movie_genres.genre_id = genres.id 
		WHERE genres.id = $1
		`
	rows, err := db.Query(query, genreID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed get data from database")
		}
		return nil, err
	}

	var movies []*domain.Movie
	for rows.Next() {
		var movie domain.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.ReleaseDate, &movie.Duration, &movie.Plot, &movie.PosterUrl, &movie.TrailerUrl, &movie.Language, &movie.CreatedAt, &movie.UpdatedAt, &genreID)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	return movies, nil
}
