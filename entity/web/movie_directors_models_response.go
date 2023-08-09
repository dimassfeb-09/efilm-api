package web

import (
	"time"
)

type Director struct {
	DirectorID  int       `json:"director_id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type MovieDirectorModelResponse struct {
	Movie     Movie      `json:"movie"`
	Directors []Director `json:"directors"`
}
