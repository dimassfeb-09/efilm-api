package web

type MovieGenreModelRequestPost struct {
	ID      int `json:"id"`
	MovieID int `json:"movie_id"`
	GenreID int `binding:"required" json:"genre_id"`
}

type MovieGenreModelRequestPut struct {
	ID      int `json:"id"`
	MovieID int `json:"movie_id"`
	GenreID int `json:"actor_id"`
}
