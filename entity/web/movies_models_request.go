package web

import "time"

type MovieModelRequest struct {
	ID          int       `json:"id"`
	Title       string    `binding:"required" json:"title"`
	ReleaseDate string    `json:"release_date"`
	Duration    int       `json:"duration"`
	Plot        string    `json:"plot"`
	PosterUrl   string    `json:"poster_url"`
	TrailerUrl  string    `json:"trailer_url"`
	Language    string    `json:"language"`
	GenreIDS    []int     `json:"genre_ids"`
	NationalID  int       `binding:"required" json:"national_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
