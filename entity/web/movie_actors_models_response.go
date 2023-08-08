package web

import (
	"github.com/dimassfeb-09/efilm-api.git/entity/domain"
	"time"
)

type MovieActorModelResponse struct {
	ID          int            `json:"id"`
	MovieID     int            `json:"movie_id`
	Title       string         `json:"title"`
	ReleaseDate time.Time      `json:"release_date"`
	Duration    int            `json:"duration"`
	Plot        string         `json:"plot"`
	PosterUrl   string         `json:"poster_url"`
	TrailerUrl  string         `json:"trailer_url"`
	Language    string         `json:"language"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Actors      []domain.Actor `json:"actors"`
}
