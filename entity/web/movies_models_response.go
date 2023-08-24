package web

import "time"

type MovieModelResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate string    `json:"release_date"`
	Duration    int       `json:"duration"`
	Plot        string    `json:"plot"`
	PosterUrl   string    `json:"poster_url"`
	TrailerUrl  string    `json:"trailer_url"`
	Language    string    `json:"language"`
	GenreIDS    []int     `json:"genre_ids"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type MoviesGenreResponse struct {
	GenreID int                   `json:"genre_id"`
	Movies  []*MovieModelResponse `json:"movies"`
}

type Movie struct {
	MovieID     int       `json:"movie_id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
}
