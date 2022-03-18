package config

import (
	"fmt"
)

type Config struct {
	Message string
	Code    uint
}

func (e Config) Error() string {
	return fmt.Sprintf("%s - %d", e.Message, e.Code)
}

func ConfigError(m string, c uint) error {
	return Config{m, c}
}
