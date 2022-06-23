package infrastructure

import (
	"log"
	"media-upload-server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	DSN := "host=" + config.C.Database.Host + " user=" + config.C.Database.User + " password=" + config.C.Database.Password + " port=" + config.C.Database.Port + " sslmode=enable" + " TimeZone=Asia/Seoul"

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
