# Backend Map

## Structure

- `internal/config` — cleanenv config struct
- `internal/http` — router, middleware, handlers, server
- `internal/service` — business logic
- `internal/repository` — sqlx data access
- `internal/model` — structs + constants
- `internal/queue` — asynq tasks + handlers
- `internal/observability` — logger, metrics, tracing

## Conventions

- Handlers thin; inject Service via constructor.
- Repositories accept context.
- Wrap errors with %w.