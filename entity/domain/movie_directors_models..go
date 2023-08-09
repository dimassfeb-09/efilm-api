package domain

type MovieDirector struct {
	Movie     Movie      `json:"movie"`
	Directors []Director `json:"directors"`
}
