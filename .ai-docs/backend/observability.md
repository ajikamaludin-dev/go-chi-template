# Observability

## What the app provides

- `GET /healthz` — liveness
- `GET /readyz` — readiness
- `GET /metrics` — Prometheus metrics (if `METRICS_ENABLED=true`)
- HTTP request/response logging to `storage/logs/<service>-app.log` (rotating)
- OpenTelemetry tracing: pushes spans via OTLP to `OTEL_ENDPOINT` (disabled if empty)
- Structured logs: `app.log`, `error.log`, `audit.log` per service prefix

## Setup external observability (optional)

### 1. OpenTelemetry Collector
Run collector with config `deploy/otel-collector-config.yaml`:
```bash
docker run -d --name otel-collector \
  -p 4317:4317 -p 4318:4318 -p 8889:8889 \
  -v $(pwd)/deploy/otel-collector-config.yaml:/etc/otel-collector-config.yaml:ro \
  otel/opentelemetry-collector-contrib:latest \
  --config=/etc/otel-collector-config.yaml
```
Set in `.env`: `OTEL_ENDPOINT=localhost:4317`.

### 2. Jaeger (traces UI)
```bash
docker run -d --name jaeger \
  -e COLLECTOR_OTLP_ENABLED=true \
  -p 16686:16686 -p 4317:4317 \
  jaegertracing/all-in-one:latest
```
UI: http://localhost:16686

### 3. Prometheus (metrics)
Create `deploy/prometheus.yml` (provided) then:
```bash
docker run -d --name prometheus \
  -p 9090:9090 \
  -v $(pwd)/deploy/prometheus.yml:/etc/prometheus/prometheus.yml:ro \
  prom/prometheus:v3.13.4
```
UI: http://localhost:9090

### 4. Grafana (dashboard)
```bash
docker run -d --name grafana \
  -p 3000:3000 \
  -e GF_SECURITY_ADMIN_PASSWORD=admin \
  -v $(pwd)/deploy/grafana/provisioning:/etc/grafana/provisioning:ro \
  grafana/grafana:13.2.4
```
UI: http://localhost:3000 (admin/admin). Add datasources Prometheus (http://host.docker.internal:9090) and Jaeger (http://host.docker.internal:16686).