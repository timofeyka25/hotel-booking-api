package server

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type Config struct {
	ReadTimeout string `env:"SERVER_READ_TIMEOUT" envDefault:"10"`
}

func FiberConfig(cfg Config) fiber.Config {
	readTimeoutSecondsCount, _ := strconv.Atoi(cfg.ReadTimeout)

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
