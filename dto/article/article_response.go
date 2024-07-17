package dto_article

import "time"

type UserResponse struct {
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Image       string    `json:"image"`
	PostedAt    time.Time `json:"posted"`
}
