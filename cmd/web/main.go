package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aji/chi-go-template/internal/config"
	httpx "github.com/aji/chi-go-template/internal/http"
	"github.com/aji/chi-go-template/internal/observability"
	"github.com/aji/chi-go-template/internal/queue"
	"github.com/aji/chi-go-template/internal/service"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load(".env")
	if err != nil {
		// try without file
		cfg, err = config.Load("")
		if err != nil {
			fmt.Println("config error:", err)
			os.Exit(1)
		}
	}

	ch, err := observability.NewLogger(cfg.Logging.Dir, cfg.Logging.Level, cfg.Logging.MaxSizeMB, cfg.Logging.MaxBackups, cfg.Logging.MaxAgeDays, "web")
	if err != nil {
		fmt.Println("logger error:", err)
		os.Exit(1)
	}

	shutdown, err := observability.InitTracing(context.Background(), cfg.App.Name, cfg.Tracing.Endpoint, cfg.Tracing.Insecure)
	if err != nil {
		ch.Error.Error("tracing init failed", zap.Error(err))
	}

	queueClient := queue.NewClient(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)
	defer queueClient.Close()

	router := httpx.NewRouter(httpx.RouterDeps{
		Cfg:        cfg,
		Logger:     ch,
		Metrics:    cfg.Metrics.Enabled,
		JobService: service.NewJobService(queueClient),
	})

	srv := httpx.NewServer(fmt.Sprintf(":%d", cfg.App.Port), router)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !httpx.IsShutdown(err) {
			ch.Error.Error("server error", zap.Error(err))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ch.App.Info("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		ch.Error.Error("shutdown error", zap.Error(err))
	}
	if shutdown != nil {
		_ = shutdown(ctx)
	}
}
