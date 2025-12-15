package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AniListClientID     string
	AniListClientSecret string
	AniListUsername     string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		AniListClientID:     os.Getenv("ANILIST_CLIENT_ID"),
		AniListClientSecret: os.Getenv("ANILIST_CLIENT_SECRET"),
		AniListUsername:     os.Getenv("ANILIST_USERNAME"),
	}

	if cfg.AniListClientID == "" || cfg.AniListClientSecret == "" || cfg.AniListUsername == "" {
		return nil, errors.New("ANILIST_CLIENT_ID, ANILIST_CLIENT_SECRET and ANILIST_USERNAME must be set in .env")
	}

	return cfg, nil
}
