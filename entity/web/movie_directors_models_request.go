package web

type MovieDirectorModelRequestPost struct {
	ID         int `json:"id"`
	MovieID    int `json:"movie_id"`
	DirectorID int `binding:"required" json:"actor_id"`
}

type MovieDirectorModelRequestPut struct {
	ID         int `json:"id"`
	MovieID    int `json:"movie_id"`
	DirectorID int `json:"actor_id"`
}
