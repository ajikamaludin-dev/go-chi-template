# Third-party Integration Pattern

Place external API clients in `internal/service/external/<provider>/`.

## Client example

```go
type Client struct {
    base string
    http *http.Client
}
func New(base string) *Client {
    return &Client{base: base, http: &http.Client{Timeout: 10*time.Second}}
}
func (c *Client) Get(ctx context.Context, path string) ([]byte, error) {
    req, _ := http.NewRequestWithContext(ctx, "GET", c.base+path, nil)
    resp, err := c.http.Do(req)
    if err != nil { return nil, err }
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}
```

## Rules

- Always use context-aware HTTP requests.
- Set timeouts.
- Wrap errors with %w.
- Retry with backoff for transient errors.
- Do not log secrets (tokens/keys).
- Store credentials in config/env, never hardcode.