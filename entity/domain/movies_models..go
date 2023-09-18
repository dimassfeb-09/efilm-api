package domain

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    int       `json:"duration"`
	Plot        string    `json:"plot"`
	PosterUrl   string    `json:"poster_url"`
	TrailerUrl  string    `json:"trailer_url"`
	Language    string    `json:"language"`
	NationalID  int       `json:"national_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
