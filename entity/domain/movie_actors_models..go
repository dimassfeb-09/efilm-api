package domain

type MovieActor struct {
	Movie  Movie   `json:"movie"`
	Actors []Actor `json:"actors"`
}
