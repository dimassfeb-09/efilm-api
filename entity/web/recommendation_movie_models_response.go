package web

type RecommendationMovieModelResponse struct {
	ID          int    `json:"movie_id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
}
