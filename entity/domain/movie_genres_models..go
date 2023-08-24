package domain

type MovieGenre struct {
	Movie    Movie `json:"movie"`
	GenreIDS []int `json:"genre_ids"`
}
