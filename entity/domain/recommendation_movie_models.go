package domain

type RecommendationMovie struct {
	ID          int    `json:"movie_id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
}
