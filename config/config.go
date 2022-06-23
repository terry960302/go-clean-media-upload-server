package config

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User                 string
		Password             string
		Net                  string
		Host                 string
		Port                 string
		DBName               string
		AllowNativePasswords bool
		Params               struct {
			ParseTime string
		}
	}
	Server struct {
		Port string
	}
}

var C config

func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// viper.AddConfigPath(filepath.Join("$GOPATH", "src", "github.com", "terry960302", "media-upload-server", "config")) // set remote file
	viper.AddConfigPath("/config.yaml") // local file
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}

	spew.Dump(C)
}
