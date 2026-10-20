# chi-go-template

Go web service template built on `chi` router, with separate web and queue worker binaries, MariaDB/Postgres/SQLite flexibility, asynq queue, and full observability.

## Stack

- Go 1.24
- chi v5
- sqlx + golang-migrate (mysql/postgres/sqlite)
- asynq (Redis) for queue/worker
- cleanenv config
- slog + lumberjack (rotating logs), Prometheus metrics, OpenTelemetry tracing

## Layout

```
cmd/web        HTTP service entrypoint
cmd/worker     queue worker entrypoint
internal/config
internal/http  router, middleware, handlers
internal/service
internal/repository
internal/model
internal/queue asynq tasks + handlers
internal/observability
migrations
```

## Using as template

1. Copy this repo to your new project.
2. Rename module path (updates go.mod + imports):

```bash
make init MODULE=github.com/your-org/your-service
# or: bash scripts/init.sh github.com/your-org/your-service
```
3. Update repo name/remote as needed. Module path must match the public repo URL if you publish.

## Quick start

```bash
cp .env.example .env
go mod download
go run ./cmd/web
# worker
go run ./cmd/worker
```

## Docker

```bash
docker compose up -d
```

## Testing

```bash
go test ./...
```

## Service access (docker compose)

| Service | URL | Credentials |
|---|---|---|
| Web API | http://localhost:8080 | - |
| Metrics | http://localhost:8080/metrics | - |
| Health | http://localhost:8080/healthz | - |
| Jaeger (traces) | http://localhost:16686 | - |
| Prometheus (metrics) | http://localhost:9090 | - |
| Grafana (dashboard) | http://localhost:3000 | admin / admin |
| OTel Collector | localhost:4317 (gRPC) | - |
| MariaDB | localhost:3306 | app / secret |
| Redis | localhost:6379 | - |
