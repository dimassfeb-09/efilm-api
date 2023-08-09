package web

import (
	"time"
)

type Actor struct {
	ActorID     int       `json:"actor_id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type MovieActorModelResponse struct {
	Movie  Movie   `json:"movie"`
	Actors []Actor `json:"actors"`
}
