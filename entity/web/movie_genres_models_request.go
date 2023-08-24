package web

type MovieGenreModelRequestPost struct {
	ID       int   `json:"id"`
	MovieID  int   `json:"movie_id"`
	GenreIDS []int `binding:"required" json:"genre_ids"`
}

type MovieGenreModelRequestPut struct {
	ID      int `json:"id"`
	MovieID int `json:"movie_id"`
	GenreID int `json:"actor_id"`
}
