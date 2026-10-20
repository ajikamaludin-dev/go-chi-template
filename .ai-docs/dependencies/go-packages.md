# Go Dependencies

## Direct

| Module | Version | Purpose |
|---|---|---|
| github.com/go-chi/chi/v5 | v5.2.1 | HTTP router |
| github.com/go-chi/cors | v1.2.1 | CORS middleware |
| github.com/golang-migrate/migrate/v4 | v4.20.1 | DB migrations |
| github.com/hibiken/asynq | v0.25.1 | Queue/worker |
| github.com/ilyakaznacheev/cleanenv | v1.5.0 | Config |
| github.com/prometheus/client_golang | v1.22.0 | Metrics |
| go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp | v0.72.0 | HTTP tracing |
| go.opentelemetry.io/otel | v1.47.0 | Tracing SDK |
| go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc | v1.35.0 | OTLP exporter |
| go.opentelemetry.io/otel/sdk | v1.47.0 | Tracing SDK |
| go.uber.org/zap | v1.28.0 | Structured logging |
| gopkg.in/natefinch/lumberjack.v2 | v2.2.1 | Log rotation |

## Checks

- Updates: `go list -m -u all`
- Vulnerabilities: install `govulncheck` (`go install golang.org/x/vuln/cmd/govulncheck@latest`) then run `govulncheck ./...`
- Verify checksums: `go mod verify`