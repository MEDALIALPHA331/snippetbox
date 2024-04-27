package config

import (
	"errors"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	PORT int
	Address string
	StaticDirPath string
}

func (c *Config) ParseConfigFromEnv() error {
	PORT, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return err
	}

	if PORT < 1024 || PORT > 49151 {
		return errors.New("Invalid Configured Port")
	}

	c.PORT = PORT

	return nil
}
