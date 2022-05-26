package db

import (
	"custom-geth/internal/errs"
	"github.com/pkg/errors"
	"time"
)

const (
	PgDriver                  = "postgres"
	defaultConnectionLifeTime = 300
)

var UnknownSQLDriverError = errors.New("unknown sql driver")

type Config struct {
	AppName            string        `mapstructure:"APPLICATION_NAME"`
	DsnMaster          string        `mapstructure:"DB_DSN_MASTER"`
	MaxOpenConnections int           `mapstructure:"MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections int           `mapstructure:"MAX_IDLE_CONNECTIONS"`
	ConnectionLifetime time.Duration `mapstructure:"CONNECTION_LIFETIME"`
	Driver             string        `mapstructure:"DB_DRIVER"`
}

func (config *Config) Validate() error {
	if config.DsnMaster == "" {
		return errors.WithMessage(errs.InvalidPropertyErr, "DB_DSN_MASTER")
	}

	if config.ConnectionLifetime == 0 {
		config.ConnectionLifetime = defaultConnectionLifeTime * time.Second
	}

	if config.Driver != "" && !IsInSlice(config.Driver, []string{PgDriver}) {
		return UnknownSQLDriverError
	} else if config.Driver == "" {
		config.Driver = PgDriver
	}

	return nil
}

func IsInSlice(search string, container []string) bool {
	for _, val := range container {
		if val == search {
			return true
		}
	}

	return false
}
