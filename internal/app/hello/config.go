package hello

import (
	"github.com/alikhil/go-server-template/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// this logger works only inside of 'hello' package
var log = utils.Logger{MinLogLevel: utils.Info}

// Config - hello app configs
type Config struct {
	ServeAddress string `default:":3333" split_words:"true"`
}

// GetConfig - returns hello app config
func GetConfig() *Config {
	var cfg = new(Config)

	var err = godotenv.Load()
	if err != nil {
		log.Warn("failed to load env: %v", err)
	}
	// you can set some prefix for all envars
	err = envconfig.Process("", cfg)
	if err != nil {
		panic(err)
	}
	return cfg

}
