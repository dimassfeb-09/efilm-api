package web

import "time"

type ActorModelRequest struct {
	ID            int       `json:"id"`
	Name          string    `json:"name" binding:"required"  example:"Lee Ji Eun"`
	DateOfBirth   string    `json:"date_of_birth" binding:"required" example:"1998-07-21"`
	NationalityID int       `json:"nationality_id" binding:"required" example:"1"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}
