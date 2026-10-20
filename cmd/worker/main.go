package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aji/chi-go-template/internal/config"
	"github.com/aji/chi-go-template/internal/observability"
	"github.com/aji/chi-go-template/internal/queue"
)

func main() {
	cfg, err := config.Load(".env")
	if err != nil {
		cfg, err = config.Load("")
		if err != nil {
			fmt.Println("config error:", err)
			os.Exit(1)
		}
	}

	ch, err := observability.NewLogger(cfg.Logging.Dir, cfg.Logging.Level, cfg.Logging.MaxSizeMB, cfg.Logging.MaxBackups, cfg.Logging.MaxAgeDays, "worker")
	if err != nil {
		fmt.Println("logger error:", err)
		os.Exit(1)
	}

	srv := queue.NewServer(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB, 10)
	mux := queue.NewMux()

	go func() {
		if err := srv.Run(mux); err != nil {
			ch.Error.Error("worker error", observability.Err(err))
			os.Exit(1)
		}
	}()

	ch.App.Info("worker started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ch.App.Info("shutting down worker")
	srv.Shutdown()
}
