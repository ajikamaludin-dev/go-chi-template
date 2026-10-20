package config

import (
	"os"
	"testing"
)

func TestLoadFromEnv(t *testing.T) {
	os.Setenv("APP_PORT", "9999")
	os.Setenv("DB_DRIVER", "sqlite")
	defer func() {
		os.Unsetenv("APP_PORT")
		os.Unsetenv("DB_DRIVER")
	}()

	cfg, err := Load("")
	if err != nil {
		t.Fatal(err)
	}
	if cfg.App.Port != 9999 {
		t.Fatalf("expected 9999, got %d", cfg.App.Port)
	}
	if cfg.Database.Driver != "sqlite" {
		t.Fatalf("expected sqlite, got %s", cfg.Database.Driver)
	}
}