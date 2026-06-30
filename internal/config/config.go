package config

import (
	"github.com/joho/godotenv"
	"github.com/caarlos0/env/v11"
)


type Config struct {
	MongoURI    string  `env:"MONGO_URI" notEmpty:"true" `
	Port  		int  `env:"PORT" notEmpty:"true" envDefault:"3000"`
}

func Load() (Config, error) {

	
	godotenv.Load()

	var cfg Config
	err := env.Parse(&cfg)

	
	if err != nil {
		return Config{}, err
	}
	
	return cfg, nil
}

