# System Overview

Go web service template: chi router HTTP service + asynq worker, MariaDB/Postgres/SQLite, Redis queue, full observability.

## Components

- `cmd/web`: HTTP service (chi)
- `cmd/worker`: asynq worker
- `internal/*`: layered modules
- Observability: slog (rotating), Prometheus, OpenTelemetry

## Data flow

Handler → Service → Repository → DB. Queue jobs via asynq client → Redis → worker.