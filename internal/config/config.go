package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App struct {
		Name string `env:"APP_NAME" env-default:"chi-go-template"`
		Env  string `env:"APP_ENV" env-default:"development"`
		Port int    `env:"APP_PORT" env-default:"8080"`
	} `env-prefix:""`

	Database struct {
		Driver   string        `env:"DB_DRIVER" env-default:"mysql"` // mysql|postgres|sqlite
		Host     string        `env:"DB_HOST" env-default:"localhost"`
		Port     int           `env:"DB_PORT" env-default:"3306"`
		Name     string        `env:"DB_NAME" env-default:"app"`
		User     string        `env:"DB_USER" env-default:"app"`
		Password string        `env:"DB_PASSWORD" env-default:"secret"`
		MaxOpen  int           `env:"DB_MAX_OPEN" env-default:"25"`
		MaxIdle  int           `env:"DB_MAX_IDLE" env-default:"5"`
		MaxLife  time.Duration `env:"DB_MAX_LIFE" env-default:"5m"`
	} `env-prefix:""`

	Redis struct {
		Addr     string `env:"REDIS_ADDR" env-default:"localhost:6379"`
		Password string `env:"REDIS_PASSWORD" env-default:""`
		DB       int    `env:"REDIS_DB" env-default:"0"`
	} `env-prefix:""`

	Logging struct {
		Level      string `env:"LOG_LEVEL" env-default:"info"`
		Dir        string `env:"LOG_DIR" env-default:"storage/logs"`
		MaxSizeMB  int    `env:"LOG_MAX_SIZE_MB" env-default:"100"`
		MaxBackups int    `env:"LOG_MAX_BACKUPS" env-default:"5"`
		MaxAgeDays int    `env:"LOG_MAX_AGE_DAYS" env-default:"30"`
	} `env-prefix:""`

	Tracing struct {
		Endpoint string `env:"OTEL_ENDPOINT" env-default:""`
		Insecure bool   `env:"OTEL_INSECURE" env-default:"true"`
	} `env-prefix:""`

	Metrics struct {
		Enabled bool   `env:"METRICS_ENABLED" env-default:"true"`
		Path    string `env:"METRICS_PATH" env-default:"/metrics"`
	} `env-prefix:""`
}

func Load(path string) (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		// fall back to env-only
		if err := cleanenv.ReadEnv(cfg); err != nil {
			return nil, fmt.Errorf("load config: %w", err)
		}
	}
	return cfg, nil
}
