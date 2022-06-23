package domain

import "time"

type VideoMetadata struct {
	ID        uint       `gorm:"primary key" json:"id"`
	MediaId   uint       `json:"mediaId"`
	Format    string     `json:"format"`
	Volume    string     `json:"volume"`
	CreatedAt *time.Time `json:"createdAt"`
}
