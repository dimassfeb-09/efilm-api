package web

type Genre struct {
	GenreID int    `json:"genre_id"`
	Name    string `json:"name"`
}

type MovieGenreModelResponse struct {
	Movie  Movie   `json:"movie"`
	Genres []Genre `json:"genres"`
}
