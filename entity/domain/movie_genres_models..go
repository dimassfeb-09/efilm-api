package domain

type MovieGenre struct {
	Movie  Movie   `json:"movie"`
	Genres []Genre `json:"directors"`
}
