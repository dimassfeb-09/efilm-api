package domain

import "time"

type ActorMovie struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Role        string    `json:"role"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type MovieActor struct {
	Movie  Movie        `json:"movie"`
	Actors []ActorMovie `json:"actors"`
}
