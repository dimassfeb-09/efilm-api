package web

import (
	"time"
)

type Movie struct {
	MovieID     int       `json:"movie_id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
}

type Director struct {
	DirectorID  int       `json:"director_id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type MovieDirectorModelResponse struct {
	Movie     Movie      `json:"movie"`
	Directors []Director `json:"directors"`
}
