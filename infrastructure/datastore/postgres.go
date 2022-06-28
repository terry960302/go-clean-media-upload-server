package infrastructure

import (
	"log"

	"github.com/terry960302/go-clean-media-upload-server/config"
	"github.com/terry960302/go-clean-media-upload-server/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresql() *gorm.DB {
	DSN := "host=" + config.C.Database.Host +
		" user=" + config.C.Database.User +
		" password=" + config.C.Database.Password +
		" port=" + config.C.Database.Port +
		" database=" + config.C.Database.DBName +
		" sslmode=disable" +
		" TimeZone=Asia/Seoul"

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.MediaMetadata{}, &domain.ImageMetadata{}, &domain.VideoMetadata{})

	return db
}
