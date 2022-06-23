package domain

import "time"

type MediaMetadata struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Url       string     `json:"url"`
	CreatedAt *time.Time `json:"created_at"`
}

// media > image, video
