package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

// Conf : the configuration objet for the app
type Conf struct {
	Server serverConf
	Debug        bool		    `env:"DEBUG,required"`
}

type serverConf struct {
	Port         int            `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration  `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration  `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration  `env:"SERVER_TIMEOUT_IDLE,required"`
	
}

// AppConfig : returns the configuration pour our app.
func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
