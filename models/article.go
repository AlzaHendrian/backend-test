package models

import "time"

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" gorm:"type: varchar(255)"`
	Description string    `json:"desc" gorm:"type: text"`
	Image       string    `json:"image" gorm:"type:varchar(255)"`
	PostedAt    string    `json:"posted"`
	Creator     string    `json:"creator" gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
