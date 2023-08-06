package web

import "time"

type DirectorModelResponse struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	NationalityID int       `json:"nationality_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
