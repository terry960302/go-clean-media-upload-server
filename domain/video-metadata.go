package domain

import "time"

type VideoMetadata struct {
	ID        uint          `gorm:"primary key" json:"id"`
	MediaID   uint          `json:"mediaId"`
	Media     MediaMetadata `gorm:"foreignKey:MediaID"`
	Format    string        `json:"format"`
	Volume    string        `json:"volume"`
	CreatedAt *time.Time    `json:"createdAt"`
}
