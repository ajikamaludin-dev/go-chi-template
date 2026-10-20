# Caching Pattern

Use Redis (already in compose) via `go-redis`. Client in config/redis.

## Example (cache-aside)

```go
func (s *Service) Get(ctx context.Context, key string) (string, error) {
    v, err := s.rdb.Get(ctx, key).Result()
    if err == nil { return v, nil }
    // load from DB
    v = loadFromDB()
    s.rdb.Set(ctx, key, v, 5*time.Minute)
    return v, nil
}
```

## Invalidation

Delete key on update: `s.rdb.Del(ctx, key)`.

## Notes

- Cache keys: prefix with domain, e.g. `user:123`.
- Avoid caching sensitive data without TTL.
- Use context timeouts.