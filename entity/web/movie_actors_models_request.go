package web

type MovieActorModelRequestPost struct {
	ID      int    `json:"id"`
	MovieID int    `json:"movie_id"`
	ActorID int    `binding:"required" json:"actor_id"`
	Role    string `binding:"required" json:"role"`
}

type MovieActorModelRequestPut struct {
	ID      int    `json:"id"`
	MovieID int    `json:"movie_id"`
	ActorID int    `json:"actor_id"`
	Role    string `binding:"required" json:"role"`
}
