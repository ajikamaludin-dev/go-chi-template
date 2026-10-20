# Testing Strategy

- Unit tests: testify, httptest.
- Integration: testcontainers-go (MariaDB/Redis) optional via env.
- Run: `go test ./...`.
- Coverage: `go test -cover ./...`.